package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rachel-lawrie/verus_backend_core/models"
	"github.com/rachel-lawrie/verus_backend_core/zaplogger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rachel-lawrie/verus_backend_core/interfaces"
)

// createVerificationLevelObject creates a new VerificationLevel object with provided name, dob, address, email, phone and auto-generates fields like VerificationLevel id and timestamps.

func ParseDocumentTypes(csv string) ([]models.DocumentType, error) {
	logger := zaplogger.GetLogger()
	var docTypes []models.DocumentType
	parts := strings.Split(csv, ",") // Split by comma

	for _, part := range parts {
		docTypeStr := strings.TrimSpace(part) // Remove extra spaces
		if docType, err := models.ParseDocumentType(docTypeStr); err != nil {
			logger.Error("Error parsing document type", zap.Error(err))
			return nil, fmt.Errorf("invalid document type: %s", docTypeStr)
		} else {
			docTypes = append(docTypes, docType)
		}
	}
	return docTypes, nil
}
func ParseOptionalGroups(input string) ([][]models.DocumentType, error) {
	logger := zaplogger.GetLogger()
	var groups [][]models.DocumentType
	groupStrings := strings.Split(input, "|") // Split by group delimiter (e.g., "|")

	for _, groupStr := range groupStrings {
		if groupStr == "" {
			continue // Skip empty groups
		}
		var docGroup []models.DocumentType
		docStrings := strings.Split(groupStr, ",") // Split by comma within a group

		for _, docStr := range docStrings {
			docTypeStr := strings.TrimSpace(docStr) // Remove spaces
			if docType, err := models.ParseDocumentType(docTypeStr); err != nil {
				logger.Error("Error parsing document type", zap.Error(err))
				return nil, fmt.Errorf("invalid document type: %s", docTypeStr)
			} else {
				docGroup = append(docGroup, docType)
			}
		}

		// Append the parsed group if it's not empty
		if len(docGroup) > 0 {
			groups = append(groups, docGroup)
		}
	}

	return groups, nil
}

func createVerificationLevelObject(name, requiredDocsStr, optionalGroupsStr string, maxAttempts int) (models.VerificationLevel, error) {
	level := models.VerificationLevel{}

	requiredDocs, err := ParseDocumentTypes(requiredDocsStr)
	if err != nil {
		return level, err
	}
	level.RequiredDocs = requiredDocs

	optionalGroups, err := ParseOptionalGroups(optionalGroupsStr)
	if err != nil {
		return level, err
	}
	level.OptionalGroups = optionalGroups

	level.Name = name
	if maxAttempts > 5 {
		return level, fmt.Errorf("maxAttempts cannot exceed 5")
	}
	level.MaxAttempts = maxAttempts
	level.LevelID = uuid.New().String()
	return level, nil
}

func CreateVerificationLevel(c *gin.Context, service interfaces.VerificationLevelService) {
	logger := zaplogger.GetLogger()
	// Define the input struct for the VerificationLevel
	var input struct {
		Name              string `json:"name" binding:"required"`        // VerificationLevel's name
		RequiredDocsStr   string `json:"requiredDocsStr"`                // VerificationLevel's requiredDocs
		OptionalGroupsStr string `json:"optionalGroupsStr"`              // VerificationLevel's optionalGroups
		MaxAttempts       int    `json:"maxAttempts" binding:"required"` // VerificationLevel's maxAttempts
	}

	// Set content type to application/json
	c.Header("Content-Type", "application/json")

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error("CreateVerificationLevel: Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	level, err := createVerificationLevelObject(input.Name, input.RequiredDocsStr, input.OptionalGroupsStr, input.MaxAttempts)
	if err != nil {
		logger.Error("Error creating VerificationLevel object", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Debug("VerificationLevel object created", zap.Any("VerificationLevel", level))

	// Call the upload service to handle the file upload
	level, err = service.CreateVerificationLevel(c, &level)
	if err != nil {
		logger.Error("Error creating VerificationLevel", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create VerificationLevel"})
		return // Exit early
	}

	c.JSON(http.StatusOK, gin.H{"message": "VerificationLevel created successfully", "VerificationLevel_id": level.LevelID})
}

// GetAllVerificationLevels is the handler function for retrieving all VerificationLevels
func GetAllVerificationLevels(c *gin.Context, service interfaces.VerificationLevelService) {
	logger := zaplogger.GetLogger()
	levels, err := service.GetAllVerificationLevels(c)
	if err != nil {
		logger.Error("GetAllVerificationLevels: Error retrieving VerificationLevels", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve VerificationLevels"})
		return
	}
	// Respond with the list of VerificationLevels
	//convert levels to a slice of maps for JSON serialization, also convert DocumentType to constant string that maps the value
	var levelsMap []map[string]interface{}
	for _, level := range levels {
		levelMap := getReadableLevel(level)
		levelsMap = append(levelsMap, levelMap)

	}
	c.JSON(http.StatusOK, levelsMap)
}

func getReadableLevel(level models.VerificationLevel) map[string]interface{} {
	requiredDocs := ""
	for j, docType := range level.RequiredDocs {
		if j > 0 {
			requiredDocs += "|"
		}
		requiredDocs += docType.String()
	}

	optionalGroups := ""
	for j, group := range level.OptionalGroups {
		if j > 0 {
			optionalGroups += "|"
		}

		groupStr := ""
		for k, docType := range group {
			if k > 0 {
				groupStr += ","
			}
			groupStr += docType.String()
		}
		optionalGroups += groupStr
	}

	levelMap := map[string]interface{}{
		"level_id":        level.LevelID,
		"name":            level.Name,
		"max_attempts":    level.MaxAttempts,
		"required_docs":   requiredDocs,
		"optional_groups": optionalGroups,
		"created_at":      level.CreatedAt,
		"updated_at":      level.UpdatedAt,
	}
	return levelMap
}

// GetDocument is the handler function for retrieving document metadata by ID
func GetVerificationLevel(c *gin.Context, service interfaces.VerificationLevelService) {
	logger := zaplogger.GetLogger()
	// Get the document ID from the URL parameter
	levelID := c.Param("id")
	logger.Debug("GetVerificationLevel: VerificationLevel ID", zap.String("id", levelID))

	level, err := service.GetVerificationLevel(c, levelID)

	if err != nil {
		// Return a JSON response with an error message if document not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	levelMap := getReadableLevel(level)
	// Respond with the document metadata
	c.JSON(http.StatusOK, levelMap)
}

// GetDocument is the handler function for retrieving document metadata by ID
func GetVerificationLevelByName(c *gin.Context, service interfaces.VerificationLevelService) {
	logger := zaplogger.GetLogger()
	// Get the document ID from the URL parameter
	levelName := c.Param("levelname")
	logger.Debug("GetVerificationLevelByName: VerificationLevel Name", zap.String("levelname", levelName))

	level, err := service.GetVerificationLevelByName(c, levelName)

	if err != nil {
		// Return a JSON response with an error message if document not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	levelMap := getReadableLevel(level)
	// Respond with the document metadata
	c.JSON(http.StatusOK, levelMap)
}

// UpdateDocument is the handler function for updating the status of a document
func UpdateVerificationLevel(c *gin.Context, service interfaces.VerificationLevelService) {
	// Get the document ID from the URL parameter
	appliantID := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		log.Printf("UpdateVerificationLevel: Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doc, err := service.UpdateVerificationLevel(c, appliantID, updates)
	if err != nil {
		// Return a JSON response with an error message if document not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated document metadata
	c.JSON(http.StatusOK, doc)
}
