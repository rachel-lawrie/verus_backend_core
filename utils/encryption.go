package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/rachel-lawrie/verus_backend_core/models"
)

// EncryptField encrypts a given plaintext using AES-GCM
func EncryptField(plaintext string, key []byte) (models.EncryptedField, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return models.EncryptedField{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return models.EncryptedField{}, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return models.EncryptedField{}, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)
	return models.EncryptedField{
		Ciphertext: ciphertext,
		Nonce:      nonce,
	}, nil
}

// EncryptAddress encrypts the address fields
func EncryptAddress(address models.RawAddress, key []byte) (models.EncryptedAddress, error) {
	line1, err := EncryptField(address.Line1, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	line2, err := EncryptField(address.Line2, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	city, err := EncryptField(address.City, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	region, err := EncryptField(address.Region, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	postalCode, err := EncryptField(address.PostalCode, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	country, err := EncryptField(address.Country, key)
	if err != nil {
		return models.EncryptedAddress{}, err
	}

	return models.EncryptedAddress{
		Line1:      line1,
		Line2:      line2,
		City:       city,
		Region:     region,
		PostalCode: postalCode,
		Country:    country,
	}, nil
}

// DecryptField decrypts a given ciphertext using AES-GCM
func DecryptField(encryptedField models.EncryptedField, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesGCM.Open(nil, encryptedField.Nonce, encryptedField.Ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// DecryptAddress decrypts the address fields
func DecryptAddress(encryptedAddress models.EncryptedAddress, key []byte) (models.RawAddress, error) {
	line1, err := DecryptField(encryptedAddress.Line1, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	line2, err := DecryptField(encryptedAddress.Line2, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	city, err := DecryptField(encryptedAddress.City, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	region, err := DecryptField(encryptedAddress.Region, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	postalCode, err := DecryptField(encryptedAddress.PostalCode, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	country, err := DecryptField(encryptedAddress.Country, key)
	if err != nil {
		return models.RawAddress{}, err
	}

	return models.RawAddress{
		Line1:      line1,
		Line2:      line2,
		City:       city,
		Region:     region,
		PostalCode: postalCode,
		Country:    country,
	}, nil
}
