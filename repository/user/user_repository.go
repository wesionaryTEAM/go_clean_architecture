package repository

import (
	"prototype2/domain"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository : get injected database
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Save(user *domain.User) (*domain.User, error) {
	return user, u.DB.Create(user).Error
}

func (u *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := u.DB.Find(&users).Error
	return users, err
}

func (u *userRepository) Delete(user *domain.User) error {
	return u.DB.Delete(&user).Error
}

func (u *userRepository) Migrate() error {
	return u.DB.AutoMigrate(&domain.User{}).Error
}
