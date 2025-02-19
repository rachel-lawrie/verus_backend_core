package services

import (
	"net/http"
	"sync"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rachel-lawrie/verus_backend_core/common"
	"github.com/rachel-lawrie/verus_backend_core/constants"
	"github.com/rachel-lawrie/verus_backend_core/models"
	"github.com/rachel-lawrie/verus_backend_core/utils"
	"github.com/rachel-lawrie/verus_backend_core/zaplogger"
	"go.mongodb.org/mongo-driver/bson"
	zap "go.uber.org/zap"
)

type VerificationLevelServiceImpl struct {
	CollectionName string
}

var (
	instance VerificationLevelServiceImpl
	once     sync.Once
)

func GetVerificationLevelServiceImpl() VerificationLevelServiceImpl {
	once.Do(func() {
		instance = VerificationLevelServiceImpl{
			CollectionName: constants.CollectionVerificationLevels,
		}
	})
	return instance
}

func (vl *VerificationLevelServiceImpl) CreateVerificationLevel(c *gin.Context, level *models.VerificationLevel) (models.VerificationLevel, error) {
	logger := zaplogger.GetLogger()
	// Get the client ID from the context
	clientIDStr, err := utils.GetClientIDFromContext(c)
	if err != nil {
		return *level, err
	}
	level.ClientID = clientIDStr

	// Check if the VerificationLevel already exists with the levelName
	existedLevel, _ := vl.GetVerificationLevelByName(c, level.Name)
	if existedLevel.Name == level.Name {
		logger.Error("VerificationLevel already exists", zap.String("Name", level.Name))
		return *level, fmt.Errorf("VerificationLevel already exists")
	}

	collection := common.GetCollection(vl.CollectionName)

	_, err = collection.InsertOne(c.Request.Context(), level)
	if err != nil {
		logger.Error("Error inserting VerificationLevel into MongoDB", zap.Error(err))
		return *level, err
	}
	return *level, nil
}

func (vl *VerificationLevelServiceImpl) GetAllVerificationLevels(c *gin.Context) ([]models.VerificationLevel, error) {
	logger := zaplogger.GetLogger()

	var levels []models.VerificationLevel

	// Get the client ID from the context
	clientIDStr, err := utils.GetClientIDFromContext(c)
	if err != nil {
		return levels, err
	}

	collection := common.GetCollection(vl.CollectionName)
	cursor, err := collection.Find(c.Request.Context(), bson.M{"client_id": clientIDStr, "deleted": false})
	if err != nil {
		logger.Error("Error fetching VerificationLevels from MongoDB", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var VerificationLevel models.VerificationLevel
		if err := cursor.Decode(&VerificationLevel); err != nil {
			logger.Error("Error decoding VerificationLevel", zap.Error(err))
			return nil, err
		}
		levels = append(levels, VerificationLevel)
	}

	if err := cursor.Err(); err != nil {
		logger.Error("Cursor error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
		return nil, err
	}

	// c.JSON(http.StatusOK, VerificationLevels)
	return levels, nil
}

func (vl *VerificationLevelServiceImpl) GetVerificationLevel(c *gin.Context, levelID string) (models.VerificationLevel, error) {
	logger := zaplogger.GetLogger()
	var level models.VerificationLevel

	// Get the client ID from the context
	clientIDStr, err := utils.GetClientIDFromContext(c)
	if err != nil {
		return level, err
	}

	// Generate the filter
	filter := bson.M{"client_id": clientIDStr, "level_id": levelID, "deleted": false}

	// Fetch the VerificationLevel directly from the database
	collection := common.GetCollection(vl.CollectionName)
	if collection == nil {
		err := fmt.Errorf("failed to get MongoDB collection: %s", vl.CollectionName)
		logger.Error("Database collection not found", zap.Error(err))
		return level, err
	}

	// Perform the database query
	err = collection.FindOne(c.Request.Context(), filter).Decode(&level)
	if err != nil {
		logger.Error("Error fetching VerificationLevel from MongoDB", zap.Error(err), zap.String("LevelID", levelID))
		return level, err
	}

	// Add a debug log to inspect the fetched VerificationLevel record
	logger.Debug("Raw VerificationLevel Record from Database", zap.Any("rawVerificationLevel", level))

	return level, nil
}

func (vl *VerificationLevelServiceImpl) GetVerificationLevelByName(c *gin.Context, levelName string) (models.VerificationLevel, error) {
	logger := zaplogger.GetLogger()
	var level models.VerificationLevel

	// Get the client ID from the context
	clientIDStr, err := utils.GetClientIDFromContext(c)
	if err != nil {
		return level, err
	}

	// Generate the filter
	filter := bson.M{"client_id": clientIDStr, "name": levelName, "deleted": false}

	// Fetch the VerificationLevel directly from the database
	collection := common.GetCollection(vl.CollectionName)
	if collection == nil {
		err := fmt.Errorf("failed to get MongoDB collection: %s", vl.CollectionName)
		logger.Error("Database collection not found", zap.Error(err))
		return level, err
	}

	// Perform the database query
	err = collection.FindOne(c.Request.Context(), filter).Decode(&level)
	if err != nil {
		logger.Error("Error fetching VerificationLevel from MongoDB", zap.Error(err), zap.String("LevelName", levelName))
		return level, err
	}

	// Add a debug log to inspect the fetched VerificationLevel record
	logger.Debug("Raw VerificationLevel Record from Database", zap.Any("rawVerificationLevel", level))
	return level, nil
}

func (vl *VerificationLevelServiceImpl) UpdateVerificationLevel(c *gin.Context, levelID string, updates map[string]interface{}) (models.VerificationLevel, error) {
	logger := zaplogger.GetLogger()
	level := models.VerificationLevel{}
	// Get the client ID from the context
	clientIDStr, err := utils.GetClientIDFromContext(c)
	if err != nil {
		return level, err
	}

	// Generate the filter and cache key
	filter, cacheKey, err := GenerateFilterAndCacheKey(levelID, clientIDStr, vl.CollectionName)
	if err != nil {
		logger.Error("Error generating filter and cache key", zap.Error(err))
		return level, err
	}

	// Build the update document
	updateDoc := bson.M{}
	for field, value := range updates {
		updateDoc[field] = value
	}
	updateDoc["updated_at"] = time.Now() // Always update the updated_at field

	update := bson.M{"$set": updateDoc}

	// Invalidate the cache after a successful update
	err = common.InvalidateCache(c, vl.CollectionName, cacheKey, filter, update, nil)
	if err != nil {
		logger.Error("Error invalidating cache", zap.Error(err))
		return level, err
	}

	// Retrieve the updated document
	result, err := vl.GetVerificationLevel(c, levelID)
	if err != nil {
		logger.Error("Error retrieving updated document", zap.Error(err))
		return level, err
	}
	return result, err
}

// GenerateFilterAndCacheKey generates the filter and cache key for a document
func GenerateFilterAndCacheKey(levelID, clientID, collectionName string) (bson.M, string, error) {
	filter := bson.M{
		"level_id":  levelID,
		"client_id": clientID,
		"deleted":   false,
	}
	cacheKey, err := common.GenerateCacheKey(collectionName, filter)
	if err != nil {
		return nil, "", err
	}
	return filter, cacheKey, nil
}
