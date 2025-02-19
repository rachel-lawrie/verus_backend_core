package models

import (
	"time"
)

// CockpitUser represents a user in the cockpit system
type CockpitUser struct {
	CockpitUserID    string     `bson:"cockpit_user_id" json:"cockpit_user_id"`       // Unique ID for the user
	Email            string     `bson:"email" json:"email"`                           // User's email address (used as username)
	Password         string     `bson:"password" json:"password"`                     // Hashed password
	Name             string     `bson:"name" json:"name"`                             // User's name
	ClientID         string     `bson:"client_id" json:"client_id"`                   // ID of the associated client (foreign key)
	CreatedAt        time.Time  `bson:"created_at" json:"created_at"`                 // Creation timestamp
	UpdatedAt        time.Time  `bson:"updated_at" json:"updated_at"`                 // Last update timestamp
	Deleted          bool       `bson:"deleted" json:"deleted"`                       // Soft delete flag
	DeletedAt        *time.Time `bson:"deleted_at" json:"deleted_at"`                 // Soft delete timestamp
	DeletedBy        *string    `bson:"deleted_by" json:"deleted_by"`                 // User/system that deleted the user
	ResetToken       *string    `bson:"reset_token" json:"reset_token"`               // Token for resetting the password (optional)
	ResetTokenExpiry *time.Time `bson:"reset_token_expiry" json:"reset_token_expiry"` // Expiry of the reset token (optional)
}
