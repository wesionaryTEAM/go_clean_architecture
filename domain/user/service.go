package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
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

// GetRawUserFromID gets the raw user from id
func (r *Repository) GetRawUserFromID(userID uint) (user *models.User, err error) {
	r.logger.Info("[UserRepository...GetRawUserFromID]")

	query := r.Model(&models.User{}).Where("id = ?", userID).First(&user)

	return user, query.Error
}
