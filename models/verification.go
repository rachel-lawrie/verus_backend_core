package models

import (
	"time"
)

type Verification struct {
	VerificationID   string              `bson:"verification_id" json:"verification_id"`     // Unique ID for the verification
	ClientID         string              `bson:"client_id" json:"client_id"`                 // ID of the associated client
	ApplicantID      string              `bson:"applicant_id" json:"applicant_id"`           // ID of the associated user
	VerficiationType VerificationType    `bson:"verification_type" json:"verification_type"` // Type of verification (e.g., KYC, AML)
	Status           VerificationStatus  `bson:"status" json:"status"`                       // Current status of the verification
	ExternalID       string              `bson:"external_id" json:"external_id"`             // External reference ID from a third-party service
	Checks           []VerificationCheck `bson:"checks" json:"checks"`                       // List of individual checks performed during the verification
	DocumentIDs      []string            `bson:"documentIDs" json:"documentIDs"`             // List of document IDs associated with the verification
	CreatedAt        time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time           `bson:"updated_at" json:"updated_at"`
	Deleted          bool                `bson:"deleted" json:"deleted"`
	DeletedAt        *time.Time          `bson:"deleted_at" json:"deleted_at"`
	DeletedBy        *string             `bson:"deleted_by" json:"deleted_by"`
}

type VerificationType int

const (
	IDVerification VerificationType = iota
	AddressVerification
	SanctionVerification
	AMLScreening
	RiskAssessment
)

type VerificationStatus int

const (
	Pending VerificationStatus = iota
	Approved
	Rejected
	Consider
)

type VerificationCheck struct {
	CheckID           string         `bson:"check_id" json:"check_id"`                     // Unique identifier for the check
	CheckType         string         `bson:"check_type" json:"check_type"`                 // Type of check (e.g., "document", "address")
	AggregationResult string         `bson:"aggregation_result" json:"aggregation_result"` // Aggregated result of the check (e.g., "passed", "failed")
	VendorResults     []VendorResult `bson:"vendor_results" json:"vendor_results"`         // Results from multiple vendors
	CreatedAt         time.Time      `bson:"created_at" json:"created_at"`                 // When the check was created
	UpdatedAt         time.Time      `bson:"updated_at" json:"updated_at"`                 // When the check was last updated
	Deleted           bool           `bson:"deleted" json:"deleted"`                       // Soft delete flag
	DeletedAt         *time.Time     `bson:"deleted_at" json:"deleted_at"`                 // Timestamp of soft deletion
	DeletedBy         *string        `bson:"deleted_by" json:"deleted_by"`                 // User or system that deleted the check
}

type VendorResult struct {
	VendorResultID string     `bson:"vendor_result_id" json:"vendor_result_id"` // Unique ID for the vendor result
	VendorID       string     `bson:"vendor_id" json:"vendor_id"`               // Foreign key: ID of the vendor providing the result
	Result         string     `bson:"result" json:"result"`                     // JSON string representing the result from the vendor
	Status         string     `bson:"status" json:"status"`                     // Status of the result (e.g., "passed", "failed"), supplied by vendor
	CreatedAt      time.Time  `bson:"created_at" json:"created_at"`             // When the check was created
	UpdatedAt      time.Time  `bson:"updated_at" json:"updated_at"`             // When the check was last updated
	Deleted        bool       `bson:"deleted" json:"deleted"`                   // Soft delete flag
	DeletedAt      *time.Time `bson:"deleted_at" json:"deleted_at"`             // Timestamp of soft deletion
	DeletedBy      *string    `bson:"deleted_by" json:"deleted_by"`             // User or system that deleted the check
}
