package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

// FindOne mocks the FindOne method of *mongo.Collection
func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	args := m.Called(ctx, filter, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

// MockSingleResult mimics *mongo.SingleResult
type MockSingleResult struct {
	mock.Mock
}

// Decode mocks the Decode method of *mongo.SingleResult
func (m *MockSingleResult) Decode(v interface{}) error {
	args := m.Called(v)
	if doc, ok := args.Get(0).(*interface{}); ok {
		*v.(*interface{}) = *doc
	}
	return args.Error(1)
}

func (m *MockSingleResult) Err() error {
	args := m.Called()
	return args.Error(0)
}

var getCollectionFunc func(name string) *MockCollection

// OverrideGetCollection allows overriding the function that gets a MongoDB collection
func OverrideGetCollection(fn func(name string) *MockCollection) {
	getCollectionFunc = fn
}

// GetCollection returns the overridden collection if set, otherwise nil
func GetCollection(name string) *MockCollection {
	if getCollectionFunc != nil {
		return getCollectionFunc(name)
	}
	return nil
}
