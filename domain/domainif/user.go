package domainif

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/types"
)

type UserService interface {
	GetOneUser(userID types.BinaryUUID) (user models.User, err error)
	UpdateUser(user *models.User) error
	DeleteUser(uuid types.BinaryUUID) error
	Create(user *models.User) error
	GetAllUser() (users *[]models.User, count int64, err error)
}
