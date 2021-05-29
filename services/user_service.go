package services

import (
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/repository"

	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(
	logger lib.Logger,
	repository repository.UserRepository,
) UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (u UserService) WithTrx(trxHandle *gorm.DB) UserService {
	u.repository = u.repository.WithTrx(trxHandle)
	return u
}

// GetOneUser gets one user
func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// UpdateUser updates the user
func (s UserService) UpdateUser(user models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(uuid lib.BinaryUUID) error {
	return s.repository.Delete(&models.User{}, uuid).Error
}

// DeleteUser deletes the user
func (s UserService) Create(user models.User) error {
	return s.repository.Create(&user).Error
}
