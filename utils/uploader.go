package utils

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/rachel-lawrie/verus_backend_core/interfaces"
	"github.com/rachel-lawrie/verus_backend_core/zaplogger"
	"go.uber.org/zap"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
}

func (u *S3Uploader) UploadFile(ctx context.Context, file multipart.File, fileName string, mimeType string, kmsUploader interfaces.KMSUploader) (string, error) { // Generate a DEK using KMSUploader

	logger := zaplogger.GetLogger()
	plaintextKey, encryptedKey, err := kmsUploader.GenerateDataKey(ctx)
	if err != nil {
		logger.Error("failed to generate data key", zap.Error(err))
		return "", fmt.Errorf("failed to generate data key: %v", err)
	}

	// Step 2: Read the file content into memory
	var fileBuffer bytes.Buffer
	_, err = io.Copy(&fileBuffer, file)
	if err != nil {
		logger.Error("failed to read file", zap.Error(err))
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	plaintextData := fileBuffer.Bytes()

	// Step 3: Encrypt the file using AES-GCM
	block, err := aes.NewCipher(plaintextKey)
	if err != nil {
		logger.Error("failed to create AES cipher", zap.Error(err))
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	nonce := make([]byte, 12) // GCM nonce size is 12 bytes
	if _, err = rand.Read(nonce); err != nil {
		logger.Error("failed to generate nonce", zap.Error(err))
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error("failed to create AES-GCM", zap.Error(err))
		return "", fmt.Errorf("failed to create AES-GCM: %v", err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintextData, nil)

	// Step 4: Upload the encrypted file to S3 with metadata
	_, err = u.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &u.BucketName,
		Key:         &fileName,
		Body:        bytes.NewReader(ciphertext),
		ContentType: &mimeType,
		ACL:         types.ObjectCannedACLPrivate,
		Metadata: map[string]string{
			"encrypted-key": base64.StdEncoding.EncodeToString(encryptedKey),
			"nonce":         base64.StdEncoding.EncodeToString(nonce),
		},
	})
	if err != nil {
		logger.Error("failed to upload encrypted file to S3", zap.Error(err))
		return "", fmt.Errorf("failed to upload encrypted file to S3: %v", err)
	}

	// Generate S3 file URL
	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.BucketName, fileName)
	return fileURL, nil
}

// This function is for testing purposes only
func (u *S3Uploader) DownloadFile(ctx context.Context, objectKey string) (*s3.GetObjectOutput, error) {
	// Call S3's GetObject API to fetch the file
	logger := zaplogger.GetLogger()
	output, err := u.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &u.BucketName,
		Key:    &objectKey,
	})
	if err != nil {
		logger.Error("failed to get object from S3", zap.Error(err))
		return nil, fmt.Errorf("failed to get object from S3: %w", err)
	}

	return output, nil
}

func NewS3Uploader(bucketName string, region string, accessKey string, secretAccessKey string) (*S3Uploader, error) {
	logger := zaplogger.GetLogger()

	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		accessKey,
		secretAccessKey,
		"",
	))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
	)
	if err != nil {
		logger.Error("failed to load AWS configuration", zap.Error(err))
		return nil, fmt.Errorf("unable to load AWS configuration: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Uploader{
		Client:     client,
		BucketName: bucketName,
	}, nil
}
