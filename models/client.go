package models

import "time"

type Client struct {
	CompanyName          string                `bson:"company_name" json:"company_name"`
	ClientID             string                `bson:"client_id" json:"client_id"`
	ClientConfigs        []ClientConfig        `bson:"client_configs" json:"client_configs"`               // embedded
	VerificationSettings []VerificationSetting `bson:"verification_settings" json:"verification_settings"` // embedded`
	Webhook              ClientWebhook         `bson:"webhook" json:"webhook"`                             // Single webhook configuration
	CreatedAt            time.Time             `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time             `bson:"updated_at" json:"updated_at"`
	Deleted              bool                  `bson:"deleted" json:"deleted"`
	DeletedAt            *time.Time            `bson:"deleted_at" json:"deleted_at"`
	DeletedBy            *string               `bson:"deleted_by" json:"deleted_by"`
}

// ClientConfig represents a configuration parameter for a client
type ClientConfig struct {
	ClientConfigID string     `bson:"client_config_id" json:"client_config_id"`
	ParameterName  string     `bson:"parameter_name" json:"parameter_name"`   // Name of the configuration parameter
	ParameterValue string     `bson:"parameter_value" json:"parameter_value"` // Value of the configuration parameter
	CreatedAt      time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `bson:"updated_at" json:"updated_at"`
	Deleted        bool       `bson:"deleted" json:"deleted"`
	DeletedAt      *time.Time `bson:"deleted_at" json:"deleted_at"` // Soft delete timestamp
	DeletedBy      *string    `bson:"deleted_by" json:"deleted_by"` // User who deleted the config
}

// VerificationSetting represents a setting for verification processes
type VerificationSetting struct {
	SettingID        string           `bson:"setting_id" json:"setting_id"`               // Unique ID for the setting
	VendorID         string           `bson:"vendor_id" json:"vendor_id"`                 // ID of the vendor providing the verification
	VerificationType VerificationType `bson:"verification_type" json:"verification_type"` // Type of verification (e.g., KYC, AML, etc.)
	ValidationCheck  string           `bson:"validation_check" json:"validation_check"`
	CreatedAt        time.Time        `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time        `bson:"updated_at" json:"updated_at"`
	Deleted          bool             `bson:"deleted" json:"deleted"`
	DeletedAt        *time.Time       `bson:"deleted_at" json:"deleted_at"`
	DeletedBy        *string          `bson:"deleted_by" json:"deleted_by"`
}

type VerificationLevel struct {
	LevelID        string           `bson:"level_id" json:"level_id"`             // Unique ID for the verification level
	ClientID       string           `bson:"client_id" json:"client_id"`           // ID of the associated client
	Name           string           `bson:"name" json:"name"`                     // Verification level name
	RequiredDocs   []DocumentType   `bson:"requiredDocs" json:"requiredDocs"`     // Mandatory docs (use DocumentType)
	OptionalGroups [][]DocumentType `bson:"optionalGroups" json:"optionalGroups"` // At least one per group
	MaxAttempts    int              `bson:"maxAttempts" json:"maxAttempts"`       // Maximum attempts allowed
	CreatedAt      time.Time        `bson:"created_at" json:"created_at"`         // Creation timestamp
	UpdatedAt      time.Time        `bson:"updated_at" json:"updated_at"`         // Last update timestamp
	Deleted        bool             `bson:"deleted" json:"deleted"`               // Soft delete flag
	DeletedAt      *time.Time       `bson:"deleted_at" json:"deleted_at"`         // Soft delete timestamp
}

// ClientWebhook represents a webhook document
type ClientWebhook struct {
	WebhookID     string      `bson:"webhook_id" json:"webhook_id"`
	ClientID      string      `bson:"client_id" json:"client_id"` // Reference to the client
	URL           string      `bson:"url" json:"url"`
	EventTypes    []EventType `bson:"event_types" json:"event_types"`
	ExpirySeconds int         `bson:"expiry_seconds" json:"expiry_seconds"` // Expiry for webhook payloads (in seconds)
	Enabled       bool        `bson:"enabled" json:"enabled"`
	SecretKey     string      `bson:"secret_key" json:"secret_key"` // This is the same as the client secret -- it's used to sign the webhook payload
	CreatedAt     time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time   `bson:"updated_at" json:"updated_at"`
	Deleted       bool        `bson:"deleted" json:"deleted"`
	DeletedAt     *time.Time  `bson:"deleted_at" json:"deleted_at"` // Soft delete timestamp
	DeletedBy     *string     `bson:"deleted_by" json:"deleted_by"` // User who deleted the config
}

type EventType string

const (
	ApplicantCreated   EventType = "applicantCreated"
	ApplicantPending   EventType = "applicantPending"
	ApplicantOnHold    EventType = "applicantOnHold"
	ApplicantReviewed  EventType = "applicantReviewed"
	ApplicantRejected  EventType = "applicantRejected"
	ApplicantDeleted   EventType = "applicantDeleted"
	InspectionReopened EventType = "inspectionReopened"
	IntegrationTest    EventType = "integrationTest"
)
