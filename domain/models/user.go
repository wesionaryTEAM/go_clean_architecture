package models

import (
	"clean-architecture/domain/constants"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	CognitoUID *string `json:"-" gorm:"index;size:50;unique"`

	FirstName   string `json:"first_name" gorm:"size:255" binding:"required"`
	LastName    string `json:"last_name" gorm:"size:255"`
	FirstNameJa string `json:"first_name_ja" gorm:"size:255"`
	LastNameJa  string `json:"last_name_ja" gorm:"size:255"`

	Email string             `json:"email" gorm:"notnull;index,unique;size:255" binding:"required,email"`
	Role  constants.UserRole `json:"role" gorm:"size:25" copier:"-"`

	IsActive        bool `json:"is_active" gorm:"default:false"`
	IsEmailVerified bool `json:"is_email_verified" gorm:"default:false"`
}

func (*User) TableName() string {
	return "users"
}

// custom validation for User model
func (u *User) Validate() error {
	return validation.Errors{
		"FirstName": validation.Validate(u.FirstName, validation.Required),
		"Email":     validation.Validate(u.Email, validation.Required, is.Email),
	}
}
