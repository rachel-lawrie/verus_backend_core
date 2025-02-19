package models

import "time"

type Secret struct {
	SecretID         string      `bson:"secret_id" json:"secret_id"`                   // Unique ID for the secret
	ClientSecretHash string      `bson:"client_secret_hash" json:"client_secret_hash"` // Hashed API key for security
	ClientID         string      `bson:"client_id" json:"client_id"`                   // References the client in the clients collection
	Name             string      `bson:"name" json:"name"`                             // Name of the secret
	IssuedAt         time.Time   `bson:"issued_at" json:"issued_at"`                   // When the secret was created
	Environment      Environment `bson:"environment" json:"environment"`               // e.g., "production" or "test"
	Revoked          bool        `bson:"revoked" json:"revoked"`                       // True if the secret has been revoked
	CreatedAt        time.Time   `bson:"created_at" json:"created_at"`
	Deleted          bool        `bson:"deleted" json:"deleted"`
	DeletedAt        *time.Time  `bson:"deleted_at" json:"deleted_at"`
	DeletedBy        *string     `bson:"deleted_by" json:"deleted_by"`
}

// Environment represents the environment for a secret (e.g., development, production, sandbox)
type Environment string

// Constants for Environment
const (
	Development Environment = "dev"     // Development environment
	Production  Environment = "prod"    // Production environment
	Sandbox     Environment = "sandbox" // Sandbox environment
)
