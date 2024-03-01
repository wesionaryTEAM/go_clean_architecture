package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/types"

	"github.com/gin-gonic/gin"
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

// WithTrx delegates transaction to repository database
func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// PaginationScope
func (s Service) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) Service {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

// GetOneUser gets one user
func (s Service) GetOneUser(userID types.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetAllUser get all the user
func (s Service) GetAllUser() (response map[string]interface{}, err error) {
	var users []models.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// UpdateUser updates the user
func (s Service) UpdateUser(user *models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s Service) DeleteUser(uuid types.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.User{}).Error
}

// DeleteUser deletes the user
func (s Service) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}
