package service

import(
	"log"
	"errors"
	"math/rand"
	"prototype2/domain"
	"time"
	"sync"
)

var once sync.Once

type userService struct {
	userRepository domain.UserRepository
}

var instance *userService

//NewUserService: construction function, injected by user repository
func NewUserService(r domain.UserRepository) domain.UserService {
	once.Do(func(){
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
		err := errors.New("The user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The name of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("The email of user is empty")
		return err
	}
	if user.DOB == "" {
		err := errors.New("The DOB of user is empty")
		return err
	}
	return nil
}

func (*userService) ValidateAge(user *domain.User) bool {
	log.Print("[UserService]...ValidateAge")
		ageLimit := 13
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	dob, err := time.Parse("2006-01-02",user.DOB)
	if err != nil {
		return false
	}

	diff := now.Sub(dob)
	diffInYears := int(diff.Hours() / (24 * 7 * 4 * 12))
	if diffInYears < ageLimit {
		return false
	} else {
		return true
	}
}

func (u *userService) Create(user *domain.User) (*domain.User, error) {
	log.Print("[UserService]...Create")
	user.ID = rand.Int63()
	return u.userRepository.Save(user)
}

func (u *userService) FindAll() ([]domain.User, error) {
	log.Print("[UserService]...FindAll")
	return u.userRepository.FindAll()
}