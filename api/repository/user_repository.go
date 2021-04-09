package repository

import (
	"clean-architecture/interfaces"
	"clean-architecture/lib"
)

type UserRepo struct {
	db lib.Database
}

func NewUserRepo(db lib.Database) interfaces.UserRepositoryInterface {
	return UserRepo{db: db}
}

func (u UserRepo) CreateUser() {

}
