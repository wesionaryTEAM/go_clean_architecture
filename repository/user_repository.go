package repository

import (
	"clean-architecture/infrastructure"
	"clean-architecture/models"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Zap.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.db.DB = trxHandle
	return r
}

// GetAll gets all users
func (r UserRepository) GetAll() (users []models.User, err error) {
	return users, r.db.Find(&users).Error
}

// Save user
func (r UserRepository) Save(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

// Update updates user
func (r UserRepository) Update(user models.User) (models.User, error) {
	return user, r.db.Save(&user).Error
}

// GetOne gets ont user
func (r UserRepository) GetOne(id uint) (user models.User, err error) {
	return user, r.db.Where("id = ?", id).First(&user).Error
}

// Delete deletes the row of data
func (r UserRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
