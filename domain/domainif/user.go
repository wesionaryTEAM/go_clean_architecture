package domainif

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/types"

	"gorm.io/gorm"
)

type UserService interface {
	GetOneUser(userID types.BinaryUUID) (user models.User, err error)
	UpdateUser(user *models.User) error
	DeleteUser(uuid types.BinaryUUID) error
	Create(user *models.User) error
	GetAllUser() (response map[string]interface{}, err error)
	SetPaginationScope(scope func(*gorm.DB) *gorm.DB) UserService
}
