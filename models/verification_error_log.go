package models

import "time"

// VerificationErrorLog represents an audit log entry for error from verification
type VerificationErrorLog struct {
	ErrorLogID     string    `bson:"error_log_id" json:"error_log_id"`       // Unique identifier for the log entry
	VerificationID string    `bson:"verification_id" json:"verification_id"` // Foreign key: ID of the associated verification
	ErrorMessage   string    `bson:"error_message" json:"error_message"`     // The error message
	ErrorCode      string    `bson:"error_code" json:"error_code"`           // Error code
	Timestamp      time.Time `bson:"timestamp" json:"timestamp"`             // Timestamp of error
}
