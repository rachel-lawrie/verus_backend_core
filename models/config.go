package models

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
