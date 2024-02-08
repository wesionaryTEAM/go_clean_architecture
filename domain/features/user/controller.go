package user

import (
	"clean-architecture/domain/domainif"
	"clean-architecture/domain/models"
	"clean-architecture/pkg/api_errors"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service domainif.UserService
	logger  framework.Logger
}

// NewUserController creates new user controller
func NewController(userService domainif.UserService, logger framework.Logger) *Controller {
	return &Controller{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u *Controller) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
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
func (u *Controller) GetUser(c *gin.Context) {
	users, err := u.service.SetPaginationScope(utils.Paginate(c)).GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}

	responses.JSONWithPagination(c, 200, users)
}

// SaveUser saves the user
func (u *Controller) SaveUser(c *gin.Context) {
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
func (u *Controller) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
	if err != nil {
		utils.HandleValidationError(u.logger, c, api_errors.ErrInvalidUUID)
		return
	}

	user, err := u.service.GetOneUser(userID)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	if err := utils.CustomBind(c.Request, &user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	metadata, _ := c.MustGet(framework.File).(types.UploadedFiles)
	user.ProfilePic = utils.SignedURL(metadata.GetFile("file").URL)

	if err := u.service.UpdateUser(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u *Controller) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
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
