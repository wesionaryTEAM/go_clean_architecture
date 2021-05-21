package repository

import (
	"clean-architecture/infrastructure"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	infrastructure.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{db}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	r.Database.DB = trxHandle
	return r
}
