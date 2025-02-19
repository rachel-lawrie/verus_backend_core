package common

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rachel-lawrie/verus_backend_core/zaplogger"

	"github.com/patrickmn/go-cache"
	"github.com/rachel-lawrie/verus_backend_core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	Client       *mongo.Client
	databaseName string       // Variable to hold the current database name
	cacheStore   *cache.Cache // In-memory cache store
)

// CollectionInterface defines the methods used from mongo.Collection
type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	// Add other methods as needed
}

// CollectionInterface defines the methods used from mongo.Collection
type SingoleResultInterface interface {
	Decode(v interface{}) error
	Err() error
	// Add other methods as needed
}

func ConnectDatabase(cfg config.DatabaseConfig) error {
	cacheStore = cache.New(time.Duration(cfg.CacheExpirationMins)*time.Minute, time.Duration(cfg.CacheCleanupIntervalMins)*time.Minute) // CacheExpirationMins-minute TTL, CacheCleanupIntervalMins-minute cleanup interval

	var mongoURI string
	if cfg.UseAtlas {
		mongoURI = cfg.AtlasConnectionURI
	} else {

		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin&authMechanism=SCRAM-SHA-256",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	}

	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("failed to create MongoDB client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB at %s:%d: %w", cfg.Host, cfg.Port, err)
	}

	// Store the database name for use in GetCollection
	databaseName = cfg.Name

	logger := zaplogger.GetLogger()
	logger.Info("Database connection established",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
	)
	return nil
}

// Helper function to simplify getting data. example: clientsCollection := GetCollection("clients")
func GetCollection(name string) *mongo.Collection {
	logger := zaplogger.GetLogger()
	// Add debug logs to inspect the state of `Client` and `databaseName`
	if Client == nil {
		logger.Fatal("MongoDB client is not initialized!",
			zap.String("function", "GetCollection"), // Log the collection name
		)
		return nil
	}
	if databaseName == "" {
		logger.Debug("Database name is not set!",
			zap.String("function", "GetCollection"), // Log the collection name
		)
		return nil
	}
	logger.Debug("Retrieving collection",
		zap.String("function", "GetCollection"), // Log the function name
		zap.String("collection", name),          // Log the collection name
		zap.String("database", databaseName),    // Log the database name
	)

	// Return the collection
	return Client.Database(databaseName).Collection(name)
}

// CacheWrapper is a helper function to fetch data from MongoDB and cache the result
func CacheWrapper(ctx context.Context, collectionName string, cacheKey string, filter interface{}, projection interface{}, result interface{}) error {
	logger := zaplogger.GetLogger()
	// Check the cache first
	cached, found := cacheStore.Get(cacheKey)
	if found {
		logger.Debug("Cache hit for key",
			zap.String("function", "CacheWrapper"),
			zap.String("cacheKey", cacheKey),
			zap.String("collection", collectionName),
		)
		// Deserialize the cached data
		jsonData, ok := cached.(string)
		if !ok {
			return fmt.Errorf("cache data format error for key: %s", cacheKey)
		}
		return json.Unmarshal([]byte(jsonData), result)
	}

	logger.Debug("Cache miss for key",
		zap.String("function", "CacheWrapper"),
		zap.String("cacheKey", cacheKey),
		zap.String("collection", collectionName),
	)
	collection := GetCollection(collectionName)
	if collection == nil {
		return fmt.Errorf("failed to get collection: %s", collectionName)
	}

	// Prepare the options with projection
	var opts *options.FindOneOptions
	if projection != nil {
		opts := options.FindOne()
		opts.SetProjection(projection)
	}

	// Use FindOne to fetch the result with the filter and projection
	singleResult := collection.FindOne(ctx, filter, opts)
	if singleResult.Err() != nil {
		return fmt.Errorf("failed to fetch data from MongoDB: %w", singleResult.Err())
	}

	// Decode the result into the provided `result` interface
	err := singleResult.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to decode MongoDB result: %w", err)
	}

	// Serialize and store the result in the cache
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to serialize data for caching: %w", err)
	}
	cacheStore.Set(cacheKey, string(jsonData), cache.DefaultExpiration)

	return nil
}

func GenerateCacheKey(collectionName string, filter interface{}) (string, error) {
	// Serialize filter and projection to JSON strings
	filterBytes, err := json.Marshal(filter)
	if err != nil {
		return "", fmt.Errorf("failed to serialize filter: %w", err)
	}

	// Generate a cache key using the collection name and serialized filter
	return fmt.Sprintf("%s:%s", collectionName, string(filterBytes)), nil
}

// UpdateCache is responsible for updating the cache after a successful update in MongoDB
func UpdateCache(ctx context.Context, collectionName string, cacheKey string, filter interface{}, update interface{}, result interface{}) error {
	logger := zaplogger.GetLogger()
	// Perform the update in the database
	collection := GetCollection(collectionName)
	if collection == nil {
		return fmt.Errorf("failed to get collection: %s", collectionName)
	}

	// Perform the update operation
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	// If no document was modified, return early
	if updateResult.ModifiedCount == 0 {
		logger.Warn("No documents were updated for filter: %v",
			zap.String("function", "UpdateCache"),
			zap.Any("filter", filter),
		)
		return nil
	}

	logger.Debug("Cache invalidation for key",
		zap.String("function", "UpdateCache"),
		zap.String("cacheKey", cacheKey),
		zap.String("collection", collectionName),
	)

	// Invalidate the old cache by deleting the cache key
	cacheStore.Delete(cacheKey)

	// Fetch the updated document and update the cache
	err = CacheWrapper(ctx, collectionName, cacheKey, filter, nil, result)
	if err != nil {
		return fmt.Errorf("failed to update cache: %w", err)
	}

	return nil
}

// Cache Invalidation - Deletes the cached entry after updating the database
func InvalidateCache(ctx context.Context, collectionName string, cacheKey string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	logger := zaplogger.GetLogger()
	// Get the MongoDB collection
	collection := GetCollection(collectionName)
	if collection == nil {
		return fmt.Errorf("failed to get collection: %s", collectionName)
	}

	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	// If no document was modified, return early
	if updateResult.ModifiedCount == 0 {
		logger.Warn("No documents were updated for filter: %v",
			zap.String("function", "InvalidateCache"),
			zap.String("collection", collectionName),
			zap.Any("filter", filter),
			zap.Any("update", update),
		)
		return nil
	}

	cacheStore.Delete(cacheKey)

	logger.Debug("Cache invalidation for key",
		zap.String("function", "InvalidateCache"),
		zap.String("cacheKey", cacheKey),
		zap.String("collection", collectionName),
	)

	return nil
}
