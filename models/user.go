package models

import "github.com/google/uuid"

// User : Database model for user
type User struct {
	ID       uuid.UUID `json:"id"`
}

func (u User) TableName() string {
	// Should return the name of table used in database. This is for the purpose of gorm
	return "users"
}
