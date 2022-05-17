package controllers

import (
	"clean-architecture/api/responses"
	"clean-architecture/api_errors"
	"clean-architecture/constants"
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/services"
	"clean-architecture/utils"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service *services.UserService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewUserController(userService *services.UserService, logger lib.Logger) *UserController {
	return &UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u *UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		utils.HandleValidationError(u.logger, c, api_errors.ErrInvalidUUID)
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u *UserController) GetUser(c *gin.Context) {
	users, err := u.service.SetPaginationScope(utils.Paginate(c)).GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}

	responses.JSONWithPagination(c, 200, users)
}

// SaveUser saves the user
func (u *UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	if err := u.service.Create(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u *UserController) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		utils.HandleValidationError(u.logger, c, api_errors.ErrInvalidUUID)
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	if err := lib.CustomBind(c.Request, &user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	metadata, _ := c.MustGet(constants.File).(lib.UploadedFiles)
	user.ProfilePic = lib.SignedURL(metadata.GetFile("file").URL)

	if err := u.service.UpdateUser(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u *UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		utils.HandleValidationError(u.logger, c, api_errors.ErrInvalidUUID)
		return
	}

	if err := u.service.DeleteUser(userID); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
