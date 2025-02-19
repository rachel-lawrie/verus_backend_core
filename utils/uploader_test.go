package utils_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rachel-lawrie/verus_backend_core/mocks"
	"github.com/rachel-lawrie/verus_backend_core/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Custom type that embeds bytes.Reader and implements multipart.File
type mockMultipartFile struct {
	*bytes.Reader
}

func (m *mockMultipartFile) Close() error {
	return nil
}

func TestUploadFile(t *testing.T) {
	mockUploader := new(mocks.MockS3Uploader)
	mockKmsClient := new(mocks.MockKMSClient)
	mockKmsUploader := &utils.KMSUploader{
		Client: mockKmsClient, // Inject the mock here
		KeyID:  "test-key-id",
	}
	fileContent := []byte("test file content")
	fileName := "testfile.txt"
	mimeType := "text/plain"

	// Create an instance of mockMultipartFile
	file := &mockMultipartFile{Reader: bytes.NewReader(fileContent)}

	mockUploader.On("UploadFile", mock.Anything, file, fileName, mimeType).Return("https://example.com/testfile.txt", nil)

	url, err := mockUploader.UploadFile(context.Background(), file, fileName, mimeType, mockKmsUploader)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/testfile.txt", url)

	mockUploader.AssertExpectations(t)
}

func TestDownloadFile(t *testing.T) {
	mockClient := new(mocks.MockS3Client)
	uploader := &mocks.MockS3Uploader{
		Client:     mockClient,
		BucketName: "test-bucket",
	}

	objectKey := "testfile.txt"
	expectedOutput := &s3.GetObjectOutput{}

	mockClient.On("GetObject", mock.Anything, &s3.GetObjectInput{
		Bucket: &uploader.BucketName,
		Key:    &objectKey,
	}).Return(expectedOutput, nil)

	mockUploader := new(mocks.MockS3Uploader)
	mockUploader.On("DownloadFile", mock.Anything, objectKey).Return(expectedOutput, nil)

	output, err := mockUploader.DownloadFile(context.Background(), objectKey)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)

	mockUploader.AssertExpectations(t)
}

func TestDownloadFileError(t *testing.T) {
	mockUploader := new(mocks.MockS3Uploader)
	objectKey := "testfile.txt"
	expectedError := assert.AnError

	mockUploader.On("DownloadFile", mock.Anything, objectKey).Return(nil, expectedError)

	output, err := mockUploader.DownloadFile(context.Background(), objectKey)
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, expectedError, err)

	mockUploader.AssertExpectations(t)
}
