package mocks

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rachel-lawrie/verus_backend_core/interfaces"
	"github.com/stretchr/testify/mock"
)

type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) GetObject(ctx context.Context, input *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}

// MockS3Uploader mocks the S3Uploader service
type MockS3Uploader struct {
	mock.Mock
	Client     *MockS3Client
	BucketName string
}

// UploadFile mocks the UploadFile method of S3Uploader
func (m *MockS3Uploader) UploadFile(ctx context.Context, file multipart.File, fileName string, mimeType string, kmsUploader interfaces.KMSUploader) (string, error) {
	args := m.Called(ctx, file, fileName, mimeType)
	return args.String(0), args.Error(1)
}

// DownloadFile mocks the DownloadFile method of S3Uploader <- This may be need changes
func (m *MockS3Uploader) DownloadFile(ctx context.Context, objectKey string) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, objectKey)
	if args.Get(0) != nil {
		return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
	}
	return nil, args.Error(1)
}
