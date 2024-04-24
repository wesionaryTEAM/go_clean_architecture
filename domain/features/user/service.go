package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/types"

	"gorm.io/gorm"
)

// UserService service layer
type Service struct {
	logger          framework.Logger
	repository      Repository
	paginationScope *gorm.DB
}

// NewUserService creates a new userservice
func NewService(
	logger framework.Logger,
	userRepository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: userRepository,
	}
}

// GetOneUser gets one user
func (s *Service) GetOneUser(userID types.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetAllUser get all the user
func (s *Service) GetAllUser() (users *[]models.User, count int64, err error) {
	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return users, count, err
	}

	return users, count, nil
}

// UpdateUser updates the user
func (s *Service) UpdateUser(user *models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s *Service) DeleteUser(uuid types.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.User{}).Error
}

// DeleteUser deletes the user
func (s *Service) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}
