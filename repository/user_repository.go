package repository

import (
	"clean-architecture/models"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	BaseRepository
}

// NewUserRepository creates a new user repository
func NewUserRepository(repo BaseRepository) UserRepository {
	return UserRepository{
		BaseRepository: repo,
	}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	r.BaseRepository = r.BaseRepository.WithTrx(trxHandle)
	return r
}

// GetAll gets all users
func (r UserRepository) GetAll() (users []models.User, err error) {
	return users, r.db.Find(&users).Error
}
