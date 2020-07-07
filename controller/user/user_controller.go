package controller

import (
	"math/rand"
	"net/http"
	"prototype2/domain"
	"prototype2/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService domain.UserService
	fbService service.FirebaseService
}

type UserController interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
}

//NewUserController: constructor, dependency injection from user service and firebase service
func NewUserController(s domain.UserService, f service.FirebaseService) UserController {
	return &userController{
		userService: s,
		fbService: f,
	}
}

func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) AddUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = rand.Int63()

	if err1 := (u.userService.Validate(&user)); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	if ageValidation := (u.userService.ValidateAge(&user)); ageValidation != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB"})
		return
	}

	password := "Zxcvbnm123"
	u.fbService.CreateUser(user.Email, password)
	u.userService.Create(&user)
	c.JSON(http.StatusOK, user)
}
