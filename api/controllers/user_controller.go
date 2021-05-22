package controllers

import (
	"clean-architecture/infrastructure"
	"clean-architecture/models"
	"clean-architecture/services"
	"clean-architecture/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  infrastructure.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger infrastructure.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	user, err := u.service.GetOneUser(models.ParseUUID(paramID))

	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		u.logger.Zap.Error(err)
	}
	c.JSON(200, gin.H{"data": users})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.Create(user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	paramID := models.ParseUUID(c.Param("id"))

	user, err := u.service.GetOneUser(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.CustomBind(c.Request, &user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("%+v \n", user)

	if err := u.service.UpdateUser(user); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	if err := u.service.DeleteUser(models.ParseUUID(paramID)); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
