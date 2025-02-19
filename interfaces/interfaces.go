package interfaces

import (
	"context"

	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	models "github.com/rachel-lawrie/verus_backend_core/models"
)

// Uploader defines the method that an uploader must implement
type Uploader interface {
	UploadFile(ctx context.Context, file multipart.File, fileName string, mimeType string, kmsUploader KMSUploader) (string, error)

	// DownloadFile downloads a file from S3 and returns the GetObjectOutput or an error
	DownloadFile(ctx context.Context, objectKey string) (*s3.GetObjectOutput, error)
}

// KMSUploader defines the methods available for KMS operations
type KMSUploader interface {
	GenerateDataKey(ctx context.Context) ([]byte, []byte, error)       // Returns plaintext and encrypted keys
	EncryptData(ctx context.Context, plaintext []byte) ([]byte, error) // Encrypts plaintext data
	DecryptData(ctx context.Context, encrypted []byte) ([]byte, error) // Decrypts encrypted data
}

type KMSClient interface {
	GenerateDataKey(ctx context.Context, input *kms.GenerateDataKeyInput, opts ...func(*kms.Options)) (*kms.GenerateDataKeyOutput, error)
	Encrypt(ctx context.Context, input *kms.EncryptInput, opts ...func(*kms.Options)) (*kms.EncryptOutput, error)
	Decrypt(ctx context.Context, input *kms.DecryptInput, opts ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

type VerificationLevelService interface {
	// UploadApplicant handles the upload of a applicant and returns metadata
	CreateVerificationLevel(c *gin.Context, verificationLevel *models.VerificationLevel) (models.VerificationLevel, error)

	// GetAllApplicants retrieves all applicants
	GetAllVerificationLevels(c *gin.Context) ([]models.VerificationLevel, error)

	// GetVerificationLevel retrieves a applicant by its ID
	GetVerificationLevel(c *gin.Context, levelID string) (models.VerificationLevel, error)

	// GetVerificationLevelByName retrieves a applicant by its name
	GetVerificationLevelByName(c *gin.Context, levelName string) (models.VerificationLevel, error)

	// UpdateApplicant updates a applicant by its ID with new data
	UpdateVerificationLevel(c *gin.Context, levelID string, updates map[string]interface{}) (models.VerificationLevel, error)
}
