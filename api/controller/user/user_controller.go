package controller

import (
	"log"
	"net/http"
	"prototype2/api/service"
	"prototype2/domain"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService domain.UserService
	fbService   service.FirebaseService
}

type UserController interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
}

//NewUserController: constructor, dependency injection from user service and firebase service
func NewUserController(s domain.UserService, f service.FirebaseService) UserController {
	return &userController{
		userService: s,
		fbService:   f,
	}
}

func (u *userController) GetUsers(c *gin.Context) {
	log.Print("[UserController]...GetUsers")
	users, err := u.userService.FindAll()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) AddUser(c *gin.Context) {
	log.Print("[UserController]...AddUser")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err1 := (u.userService.Validate(&user)); err1 != nil {
		sentry.CaptureException(err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	if ageValidation := (u.userService.ValidateAge(&user)); !ageValidation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB"})
		return
	}

	uid, err := u.fbService.CreateUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnt create user in firebase"})
		return
	}

	user.ID = uid

	userCreated, err := u.userService.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnt create user"})
		return
	}

	c.JSON(http.StatusOK, userCreated)
}
