package models

import "time"

// Vendor represents a vendor configuration and its supported features
type Vendor struct {
	VendorID          string             `bson:"vendor_id" json:"vendor_id"`                   // Unique identifier for the vendor
	VendorName        string             `bson:"vendor_name" json:"vendor_name"`               // Name of the vendor
	APIBaseURL        string             `bson:"api_base_url" json:"api_base_url"`             // Base URL for the vendor's API
	APIKey            string             `bson:"api_key" json:"api_key"`                       // API key for authenticating with the vendor
	SecretKey         string             `bson:"secret_key" json:"secret_key"`                 // API secret for authenticating with the vendor
	SupportedFeatures []string           `bson:"supported_features" json:"supported_features"` // List of supported features
	Features          []VendorFeature    `bson:"vendor_features" json:"vendor_features"`       // List of supported features
	AuthConfigs       []VendorAuthConfig `bson:"auth_configs" json:"auth_configs"`             // List of authorization configurations
	Configs           []VendorConfig     `bson:"configs" json:"configs"`                       // List of vendor-specific configurations
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`                 // Timestamp of creation
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`                 // Timestamp of last update
	Deleted           bool               `bson:"deleted" json:"deleted"`                       // Soft delete flag
	DeletedAt         *time.Time         `bson:"deleted_at" json:"deleted_at"`                 // Soft delete timestamp
	DeletedBy         *string            `bson:"deleted_by" json:"deleted_by"`                 // User/system that deleted the vendor
}

// VendorFeature represents a feature supported by a vendor with API examples
type VendorFeature struct {
	VendorFeatureID string     `bson:"vendor_feature_id" json:"vendor_feature_id"` // Unique ID for the vendor feature
	VendorID        string     `bson:"vendor_id" json:"vendor_id"`                 // Foreign key: ID of the associated vendor
	Endpoint        string     `bson:"endpoint" json:"endpoint"`                   // API endpoint for the feature
	SampleInput     string     `bson:"sample_input" json:"sample_input"`           // Example of an API request for the feature
	SampleOutput    string     `bson:"sample_output" json:"sample_output"`         // Example of an API response for the feature
	CreatedAt       time.Time  `bson:"created_at" json:"created_at"`               // Timestamp of creation
	UpdatedAt       time.Time  `bson:"updated_at" json:"updated_at"`               // Timestamp of last update
	Deleted         bool       `bson:"deleted" json:"deleted"`                     // Soft delete flag
	DeletedAt       *time.Time `bson:"deleted_at" json:"deleted_at"`               // Soft delete timestamp
	DeletedBy       *string    `bson:"deleted_by" json:"deleted_by"`               // User/system that deleted the feature
}

// VendorAuthConfig represents the authentication configuration for a vendor
type VendorAuthConfig struct {
	VendorAuthConfigID string     `bson:"vendor_auth_config_id" json:"vendor_auth_config_id"` // Unique ID for the vendor authentication config
	AuthURL            string     `bson:"auth_url" json:"auth_url"`                           // Authentication endpoint URL
	ParameterName      string     `bson:"parameter_name" json:"parameter_name"`               // Name of the parameter used in the auth request
	ParameterValue     string     `bson:"parameter_value" json:"parameter_value"`             // Value of the parameter used in the auth request
	CreatedAt          time.Time  `bson:"created_at" json:"created_at"`                       // Timestamp of creation
	UpdatedAt          time.Time  `bson:"updated_at" json:"updated_at"`                       // Timestamp of last update
	Deleted            bool       `bson:"deleted" json:"deleted"`                             // Soft delete flag
	DeletedAt          *time.Time `bson:"deleted_at" json:"deleted_at"`                       // Soft delete timestamp
	DeletedBy          *string    `bson:"deleted_by" json:"deleted_by"`                       // User/system that deleted the configuration
}

// VendorAuthConfig represents the authentication configuration for a vendor
type VendorConfig struct {
	VendorConfigID string     `bson:"vendor_config_id" json:"vendor_config_id"` // Unique ID for the vendor authentication config
	AuthURL        string     `bson:"auth_url" json:"auth_url"`                 // Authentication endpoint URL
	ParameterName  string     `bson:"parameter_name" json:"parameter_name"`     // Name of the parameter used in the auth request
	ParameterValue string     `bson:"parameter_value" json:"parameter_value"`   // Value of the parameter used in the auth request
	CreatedAt      time.Time  `bson:"created_at" json:"created_at"`             // Timestamp of creation
	UpdatedAt      time.Time  `bson:"updated_at" json:"updated_at"`             // Timestamp of last update
	Deleted        bool       `bson:"deleted" json:"deleted"`                   // Soft delete flag
	DeletedAt      *time.Time `bson:"deleted_at" json:"deleted_at"`             // Soft delete timestamp
	DeletedBy      *string    `bson:"deleted_by" json:"deleted_by"`             // User/system that deleted the configuration
}
