package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	AWS      AWSConfig
	Vendors  map[string]VendorConfig
}

type VendorConfig struct {
	WebhookSecretKey string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host                     string
	Port                     int
	User                     string
	Password                 string
	Name                     string
	CacheExpirationMins      int
	CacheCleanupIntervalMins int
	UseAtlas                 bool   // Indicate whether to use MongoDB Atlas
	AtlasConnectionURI       string // Full connection string for MongoDB Atlas
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	BucketName      string
	KeyID           string
}

func LoadConfig(env string) Config {
	viper.SetConfigName(env) // e.g., dev, sandbox, prod
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Panicf("Unable to decode into struct: %v", err)
	}

	// Load sensitive information from .env file for development
	if env == "dev" {
		if accessKey := os.Getenv("AWS_ACCESS_KEY_ID"); accessKey != "" {
			config.AWS.AccessKeyID = accessKey
		}
		if secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY"); secretKey != "" {
			config.AWS.SecretAccessKey = secretKey
		}
		if dbUsername := os.Getenv("DB_USERNAME"); dbUsername != "" {
			config.Database.User = dbUsername
		}
		if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
			config.Database.Password = dbPassword
		}
		if webhookSecret := os.Getenv("WEBHOOK_SECRET_KEY"); webhookSecret != "" {
			vendorConfig := config.Vendors["sumsub"]
			vendorConfig.WebhookSecretKey = webhookSecret
			config.Vendors["sumsub"] = vendorConfig
		}
		if awsKeyID := os.Getenv("AWS_KEY_ID"); awsKeyID != "" {
			config.AWS.KeyID = awsKeyID
		}
	}

	// Load sensitive information from secrets manager for sandbox
	if env == "sandbox" {
		secrets, err := LoadSecretsFromAWS("your-secret-name")
		if err == nil {
			config.AWS.AccessKeyID = secrets["AWS_ACCESS_KEY_ID"]
			config.AWS.SecretAccessKey = secrets["AWS_SECRET_ACCESS_KEY"]
			vendorConfig := config.Vendors["sumsub"]
			vendorConfig.WebhookSecretKey = secrets["WEBHOOK_SECRET_KEY"]
			config.Vendors["sumsub"] = vendorConfig
			config.AWS.KeyID = secrets["AWS_KEY_ID"]
		}
	}

	return config
}

// LoadSecretsFromAWS loads secrets from AWS Secrets Manager
func LoadSecretsFromAWS(secretName string) (map[string]string, error) {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(cfg)

	// Get secret value
	result, err := svc.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get secret value: %v", err)
	}

	// Parse JSON secret string into map
	var secretMap map[string]string
	if err := json.Unmarshal([]byte(*result.SecretString), &secretMap); err != nil {
		return nil, fmt.Errorf("unable to parse secret JSON: %v", err)
	}

	return secretMap, nil
}
