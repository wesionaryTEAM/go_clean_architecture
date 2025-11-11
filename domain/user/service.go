package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
)

// UserService service layer
type Service struct {
	logger     framework.Logger
	repository Repository
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

// Create creates the user in database
func (s Service) Create(user *models.User) error {
	return s.repository.Create(user).Error
}

// GetOneUser gets one user
func (s Service) GetUserByID(userID uint) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

func (r *Service) ExistsByEmail(email string) (bool, error) {
	return r.repository.ExistsByEmail(email)
}

func (s Service) GetUsers(pagination utils.Pagination) (users []models.User, count int64, err error) {
	s.logger.Info("[UserService...GetUsers]")

	return s.repository.GetAllUsers(pagination)
}
