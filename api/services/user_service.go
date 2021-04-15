package services

import (
	"clean-architecture/api/repository"
	"clean-architecture/lib"
)

type UserService struct {
	userRepo repository.UserRepo
	logger   lib.Logger
}

func NewUserService(userRepo repository.UserRepo, logger lib.Logger) UserService {
	return UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u UserService) CreateUser() {
	// Dummy Implementation
	u.userRepo.CreateUser()
}
