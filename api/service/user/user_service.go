package service

import (
	"errors"
	"log"
	"prototype2/domain"
	"sync"
	"time"
)

var once sync.Once

type userService struct {
	userRepository domain.UserRepository
}

var instance *userService

//NewUserService: construction function, injected by user repository
func NewUserService(r domain.UserRepository) domain.UserService {
	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})
	return instance
	// return &userService{
	// 	userRepository: r,
	// }
}

func (*userService) Validate(user *domain.User) error {
	log.Print("[UserService]...Validate")
	if user == nil {
		err := errors.New("the user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("the name field of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("the email field of user is empty")
		return err
	}
	if user.DOB == "" {
		err := errors.New("the DOB field of user is empty")
		return err
	}
	return nil
}

func (*userService) ValidateAge(user *domain.User) bool {
	log.Print("[UserService]...ValidateAge")
	ageLimit := 13
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return false
	}

	diff := now.Sub(dob)
	diffInYears := int(diff.Hours() / (24 * 7 * 4 * 12))
	return diffInYears >= ageLimit
}

func (u *userService) Create(user *domain.User) (*domain.User, error) {
	return u.userRepository.Save(user)
}

func (u *userService) FindAll() ([]domain.User, error) {
	log.Print("[UserService]...FindAll")
	return u.userRepository.FindAll()
}
