package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/rachel-lawrie/verus_backend_core/interfaces"
	"github.com/rachel-lawrie/verus_backend_core/zaplogger"
	"go.uber.org/zap"
)

type KMSUploader struct {
	Client interfaces.KMSClient
	KeyID  string // KMS Key ID (ARN or alias) used for generating data keys
}

func (k *KMSUploader) GenerateDataKey(ctx context.Context) ([]byte, []byte, error) {
	logger := zaplogger.GetLogger()
	result, err := k.Client.GenerateDataKey(ctx, &kms.GenerateDataKeyInput{
		KeyId:   &k.KeyID,
		KeySpec: "AES_256",
	})
	if err != nil {
		logger.Error("failed to generate data key", zap.Error(err))
		return nil, nil, err
	}
	return result.Plaintext, result.CiphertextBlob, nil
}

func (k *KMSUploader) EncryptData(ctx context.Context, plaintext []byte) ([]byte, error) {
	logger := zaplogger.GetLogger()
	result, err := k.Client.Encrypt(ctx, &kms.EncryptInput{
		KeyId:     &k.KeyID,
		Plaintext: plaintext,
	})
	if err != nil {
		logger.Error("failed to encrypt data", zap.Error(err))
		return nil, err
	}
	return result.CiphertextBlob, nil
}

func (k *KMSUploader) DecryptData(ctx context.Context, encrypted []byte) ([]byte, error) {
	logger := zaplogger.GetLogger()
	result, err := k.Client.Decrypt(ctx, &kms.DecryptInput{
		CiphertextBlob: encrypted,
	})
	if err != nil {
		logger.Error("failed to decrypt data", zap.Error(err))
		return nil, err
	}
	return result.Plaintext, nil
}

func NewKMSUploader(region, accessKey, secretAccessKey, keyID string) (*KMSUploader, error) {
	logger := zaplogger.GetLogger()
	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		accessKey,
		secretAccessKey,
		"",
	))

	// Step 2: Load AWS config with region and credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
	)
	if err != nil {
		logger.Error("failed to load AWS configuration", zap.Error(err))
		return nil, fmt.Errorf("unable to load AWS configuration: %v", err)
	}

	kmsClient := kms.NewFromConfig(cfg)
	return &KMSUploader{
		Client: kmsClient,
		KeyID:  keyID,
	}, nil
}
