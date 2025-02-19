package utils

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/rachel-lawrie/verus_backend_core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateDataKey(t *testing.T) {
	mockKmsClient := new(mocks.MockKMSClient)
	keyID := "test-key-id"

	plaintextKey := []byte("plain-key")
	ciphertextBlob := []byte("cipher-key")

	mockKmsClient.On("GenerateDataKey", mock.Anything, &kms.GenerateDataKeyInput{
		KeyId:   &keyID,
		KeySpec: "AES_256",
	}).Return(&kms.GenerateDataKeyOutput{
		Plaintext:      plaintextKey,
		CiphertextBlob: ciphertextBlob,
	}, nil)

	uploader := &KMSUploader{
		Client: mockKmsClient,
		KeyID:  keyID,
	}

	plain, cipher, err := uploader.GenerateDataKey(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, plaintextKey, plain)
	assert.Equal(t, ciphertextBlob, cipher)
	mockKmsClient.AssertExpectations(t)
}

func TestEncryptData(t *testing.T) {
	mockKmsClient := new(mocks.MockKMSClient)
	keyID := "test-key-id"
	plaintext := []byte("test-data")
	ciphertext := []byte("encrypted-data")

	mockKmsClient.On("Encrypt", mock.Anything, &kms.EncryptInput{
		KeyId:     &keyID,
		Plaintext: plaintext,
	}).Return(&kms.EncryptOutput{
		CiphertextBlob: ciphertext,
	}, nil)

	uploader := &KMSUploader{
		Client: mockKmsClient,
		KeyID:  keyID,
	}

	result, err := uploader.EncryptData(context.TODO(), plaintext)
	assert.NoError(t, err)
	assert.Equal(t, ciphertext, result)
	mockKmsClient.AssertExpectations(t)
}

func TestDecryptData(t *testing.T) {
	mockKmsClient := new(mocks.MockKMSClient)
	encryptedData := []byte("encrypted-data")
	decryptedData := []byte("decrypted-data")

	mockKmsClient.On("Decrypt", mock.Anything, &kms.DecryptInput{
		CiphertextBlob: encryptedData,
	}).Return(&kms.DecryptOutput{
		Plaintext: decryptedData,
	}, nil)

	uploader := &KMSUploader{
		Client: mockKmsClient,
		KeyID:  "test-key-id",
	}

	result, err := uploader.DecryptData(context.TODO(), encryptedData)
	assert.NoError(t, err)
	assert.Equal(t, decryptedData, result)
	mockKmsClient.AssertExpectations(t)
}
