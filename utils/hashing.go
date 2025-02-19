package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// HashAPIKey hashes the given API key using SHA-256
func HashAPIKey(apiKey string) string {
	hasher := sha256.New()
	hasher.Write([]byte(apiKey))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashPassword hashes the given password using SHA-256
func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateHMAC generates an HMAC-SHA256 hash for the given message and secret
func GenerateHMAC(message, secret string) string {
	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write([]byte(message))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateRandomString generates a secure random string of the specified length
func GenerateRandomString(length int) string {
	if length <= 0 {
		log.Fatalf("Invalid length for random string: %d", length)
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}

	// Convert bytes to a hexadecimal string
	return hex.EncodeToString(bytes)[:length]
}
