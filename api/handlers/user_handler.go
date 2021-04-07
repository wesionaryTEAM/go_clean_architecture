package handlers

import (
	"github.com/gin-gonic/gin"
	"prototype2/interfaces"
	"prototype2/lib"
)

type UserController struct {
	logger lib.Logger
	userService interfaces.UserServiceInterface
}

func NewUserController(logger lib.Logger, serviceInterface interfaces.UserServiceInterface) UserController{
	return UserController{logger: logger, userService: serviceInterface}
}

func (u *UserController) CreateUser(c * gin.Context) {

}

func (u *UserController) GetUsers(c * gin.Context) {

}

func (u *UserController) GetSingleUser(c * gin.Context) {

}

func (u *UserController) PatchUsers(c * gin.Context) {

}

func (u *UserController) DeleteUser(c * gin.Context) {

}