package services

import (
	"clean-architecture/interfaces"
	"clean-architecture/lib"
)

type UserService struct {
	userRepo interfaces.UserRepositoryInterface
	logger   lib.Logger
}

func NewUserService(userRepo interfaces.UserRepositoryInterface, logger lib.Logger) interfaces.UserServiceInterface {
	return UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u UserService) CreateUser() {
	// Dummy Implementation
	u.userRepo.CreateUser()
}
