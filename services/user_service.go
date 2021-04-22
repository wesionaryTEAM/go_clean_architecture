package services

import (
	"clean-architecture/infrastructure"
	"clean-architecture/models"
	"clean-architecture/repository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger     infrastructure.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(
	logger infrastructure.Logger,
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
func (s UserService) GetOneUser(id uint) (user models.User, err error) {
	return user, s.repository.GetOne(&user, id)
}

// GetAllUser get all the user
func (s UserService) GetAllUser() ([]models.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

// UpdateUser updates the user
func (s UserService) UpdateUser(id uint, user models.User) error {

	userDB, err := s.GetOneUser(id)
	if err != nil {
		return err
	}

	copier.Copy(&userDB, &user)

	userDB.ID = id

	return s.repository.Update(&userDB)
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id)
}

// DeleteUser deletes the user
func (s UserService) Create(user models.User) error {
	return s.repository.Create(&user)
}
