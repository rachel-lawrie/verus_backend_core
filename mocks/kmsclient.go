package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/stretchr/testify/mock"
)

// MockKMSClient mocks the AWS KMS client
type MockKMSClient struct {
	mock.Mock
}

// GenerateDataKey mocks the GenerateDataKey method of kms.Client
func (m *MockKMSClient) GenerateDataKey(ctx context.Context, input *kms.GenerateDataKeyInput, opts ...func(*kms.Options)) (*kms.GenerateDataKeyOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*kms.GenerateDataKeyOutput), args.Error(1)
}

// Encrypt mocks the Encrypt method of kms.Client
func (m *MockKMSClient) Encrypt(ctx context.Context, input *kms.EncryptInput, opts ...func(*kms.Options)) (*kms.EncryptOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*kms.EncryptOutput), args.Error(1)
}

// Decrypt mocks the Decrypt method of kms.Client
func (m *MockKMSClient) Decrypt(ctx context.Context, input *kms.DecryptInput, opts ...func(*kms.Options)) (*kms.DecryptOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*kms.DecryptOutput), args.Error(1)
}
