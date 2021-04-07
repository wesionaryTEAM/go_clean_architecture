package repository

import (
	"prototype2/interfaces"
	"prototype2/lib"
)

type UserRepo struct {
	db lib.Database
}

func NewUserRepo(db lib.Database) interfaces.UserRepositoryInterface{
	return UserRepo{db: db}
}

func (u UserRepo) CreateUser() {

}
