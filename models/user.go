package models

import (
	"clean-architecture/lib"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	lib.ModelBase
	Name         string         `json:"name" form:"name"`
	Email        string         `json:"email" form:"email"`
	Age          int            `json:"age" form:"age"`
	Birthday     *time.Time     `json:"time"`
	MemberNumber sql.NullString `json:"member_number"`
	ProfilePic   lib.SignedURL  `json:"profile_pic"`
	CreatedAt    time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" form:"updated_at"`
}

// BeforeCreate run this before creating user
func (t *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
