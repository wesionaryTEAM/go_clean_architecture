package handlers

import (
	"prototype2/interfaces"
	"prototype2/lib"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger      lib.Logger
	userService interfaces.UserServiceInterface
}

func NewUserController(logger lib.Logger, serviceInterface interfaces.UserServiceInterface) UserController {
	return UserController{logger: logger, userService: serviceInterface}
}

func (u *UserController) CreateUser(c *gin.Context) {

}

func (u *UserController) GetUsers(c *gin.Context) {

}

func (u *UserController) GetSingleUser(c *gin.Context) {

}

func (u *UserController) PatchUsers(c *gin.Context) {

}

func (u *UserController) DeleteUser(c *gin.Context) {

}
