package handlers

import (
	"clean-architecture/api/services"
	"clean-architecture/lib"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger      lib.Logger
	userService services.UserService
}

func NewUserController(logger lib.Logger, service services.UserService) UserController {
	return UserController{logger: logger, userService: service}
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
