package repository

import (
	"clean-architecture/lib"
)

type UserRepo struct {
	db lib.Database
}

func NewUserRepo(db lib.Database) UserRepo {
	return UserRepo{db: db}
}

func (u UserRepo) CreateUser() {

}
