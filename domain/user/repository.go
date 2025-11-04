package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/utils"
)

// UserRepository database structure
type Repository struct {
	*infrastructure.Database
	logger framework.Logger
}

// NewUserRepository creates a new user repository
func NewRepository(db *infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

// ExistsByEmail checks if the user exists by email
func (r *Repository) ExistsByEmail(email string) (bool, error) {
	r.logger.Info("[UserRepository...Exists]")

	users := make([]models.User, 0, 1)
	query := r.DB.Where("email = ?", email).Limit(1).Find(&users)

	return query.RowsAffected > 0, query.Error
}

func (r *Repository) GetAllUsers(pagination utils.Pagination) (users []models.User, count int64, err error) {
	r.logger.Info("[UserRepository...GetAllUsers]")

	query := r.Model(&models.User{}).Count(&count).Limit(pagination.Limit).Offset(pagination.Offset).Find(&users)

	return users, count, query.Error
}
