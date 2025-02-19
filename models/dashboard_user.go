package models

import "time"

type DashboardUser struct {
	DashboardUser string     `bson:"dashboard_user_id" json:"dashboard_user_id"`
	Name          string     `bson:"name" json:"name"`
	Email         string     `bson:"email" json:"email"`
	Phone         string     `bson:"phone" json:"phone"`
	ClientID      string     `bson:"client_id" json:"client_id"`
	ClientName    string     `bson:"client_name" json:"client_name"`
	CreatedAt     time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `bson:"updated_at" json:"updated_at"`
	Deleted       bool       `bson:"deleted" json:"deleted"`
	DeletedAt     *time.Time `bson:"deleted_at" json:"deleted_at"`
	DeletedBy     *string    `bson:"deleted_by" json:"deleted_by"`
}
