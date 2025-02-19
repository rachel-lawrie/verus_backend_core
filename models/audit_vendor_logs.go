package models

import "time"

// AuditVendorLog represents an audit log entry for actions performed with a vendor
type AuditVendorLog struct {
	LogID           string    `bson:"log_id" json:"log_id"`                     // Unique identifier for the log entry
	VendorID        string    `bson:"vendor_id" json:"vendor_id"`               // Foreign key: ID of the associated applicant
	ActionPerformed string    `bson:"action_performed" json:"action_performed"` // The action performed (e.g., "created", "updated", "deleted")
	Details         string    `bson:"details" json:"details"`                   // Additional details about the action
	Timestamp       time.Time `bson:"timestamp" json:"timestamp"`               // Timestamp of when the action occurred
}
