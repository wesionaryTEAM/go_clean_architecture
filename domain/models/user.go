package models

import (
	"clean-architecture/domain/constants"
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	UUID            types.BinaryUUID   `json:"uuid" gorm:"index;notnull"`
	CognitoUID      *string            `json:"-" gorm:"index;size:50"`
	FirstName       string             `json:"first_name" gorm:"size:255"`
	LastName        string             `json:"last_name" gorm:"size:255"`
	FirstNameJa     string             `json:"first_name_ja" gorm:"size:255"`
	LastNameJa      string             `json:"last_name_ja" gorm:"size:255"`
	Email           string             `json:"email" gorm:"notnull;index,unique;size:255"`
	Role            constants.UserRole `json:"role" gorm:"size:25" copier:"-"`
	IsActive        bool               `json:"is_active" gorm:"default:false"`
	IsEmailVerified bool               `json:"is_email_verified" gorm:"default:false"`
	IsAdmin         bool               `json:"-" gorm:"default:false"`
	ProfilePic      utils.SignedURL    `json:"profile_pic"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.UUID.String() == (types.BinaryUUID{}).String() {
		id, err := uuid.NewRandom()
		u.UUID = types.BinaryUUID(id)
		return err
	}
	return nil
}

func (*User) TableName() string {
	return "users"
}
