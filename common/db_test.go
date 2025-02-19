package common

import (
	"context"
	"testing"
	"time"

	"github.com/rachel-lawrie/verus_backend_core/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func TestConnectDatabase(t *testing.T) {
	cfg := config.DatabaseConfig{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     27017,
		Name:     "testdb",
	}

	err := ConnectDatabase(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, Client)
	assert.Equal(t, "testdb", databaseName)

	// Disconnect after test
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Disconnect(ctx)
	assert.NoError(t, err)
}

func TestGetCollection(t *testing.T) {
	cfg := config.DatabaseConfig{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     27017,
		Name:     "testdb",
	}

	err := ConnectDatabase(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, Client)
	assert.Equal(t, "testdb", databaseName)

	collection := GetCollection("testcollection")
	assert.NotNil(t, collection)
	assert.Equal(t, "testcollection", collection.Name())

	// Disconnect after test
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Disconnect(ctx)
	assert.NoError(t, err)
}

func TestGetCollectionWithoutInitialization(t *testing.T) {
	// Ensure Client and databaseName are not set
	Client = nil
	databaseName = ""

	collection := GetCollection("testcollection")
	assert.Nil(t, collection)
}
