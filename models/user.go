package models

import "github.com/google/uuid"

// User represents the user for this application
//
// User is the database model for user.
//
// swagger:model
type User struct {
	ID uuid.UUID `json:"id"`
}

func (u User) TableName() string {
	// Should return the name of table used in database. This is for the purpose of gorm
	return "users"
}
