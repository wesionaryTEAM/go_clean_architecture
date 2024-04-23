package models

import (
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	types.ModelBase
	Name         string          `json:"name" form:"name" binding:"required"`
	Email        string          `json:"email" form:"email "binding:"required,email"`
	Age          int             `json:"age" form:"age" binding:"required"`
	Birthday     *time.Time      `json:"time" form:"time" binding:"required" time_format:"2006-01-02"`
	MemberNumber sql.NullString  `json:"member_number" binding:"required"`
	ProfilePic   utils.SignedURL `json:"profile_pic"`
	CreatedAt    time.Time       `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" form:"updated_at"`
}

// BeforeCreate run this before creating user
func (t *User) BeforeCreate(_ *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = types.BinaryUUID(id)
	return err
}
