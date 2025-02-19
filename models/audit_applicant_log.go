package models

import "time"

// AuditApplicantLog represents an audit log entry for actions performed on an applicant
type AuditApplicantLog struct {
	LogID           string    `bson:"log_id" json:"log_id"`                     // Unique identifier for the log entry
	ApplicantID     string    `bson:"applicant_id" json:"applicant_id"`         // Foreign key: ID of the associated applicant
	ActionPerformed string    `bson:"action_performed" json:"action_performed"` // The action performed (e.g., "created", "updated", "deleted")
	Details         string    `bson:"details" json:"details"`                   // Additional details about the action
	Timestamp       time.Time `bson:"timestamp" json:"timestamp"`               // Timestamp of when the action occurred
	ClientID        string    `bson:"client_id" json:"client_id"`               // Foreign key: ID of the associated client
	IP              string    `bson:"ip" json:"ip"`                             // IP address from which the action was performed
}
