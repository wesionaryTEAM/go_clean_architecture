package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/api_errors"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/services"
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service         *Service
	logger          framework.Logger
	s3BucketService services.S3Service
	sesService      services.SESService
	env             *framework.Env
}

type URLObject struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewUserController creates new user controller
func NewController(
	userService *Service,
	logger framework.Logger,
	s3BucketService services.S3Service,
	sesService services.SESService,
	env *framework.Env,
) *Controller {
	return &Controller{
		service:         userService,
		logger:          logger,
		s3BucketService: s3BucketService,
		sesService:      sesService,
		env:             env,
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

func (c *Controller) UploadImage(ctx *gin.Context) {
	metadata, _ := ctx.MustGet(framework.File).(types.UploadedFiles)

	urls := make([]URLObject, 0, 1)

	for _, file := range metadata.GetMultipleFiles("file") {
		signedUrl, err := c.s3BucketService.GetSignedURL(file.URL)
		if err != nil {
			c.logger.Info("Failed to get the presigned url: ", err.Error())
			return
		}
		urls = append(urls, URLObject{Name: file.FileName, URL: signedUrl})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"urls":    urls,
	})
}

func (c *Controller) SendEmail(ctx *gin.Context) {
	var emailData services.EmailParams

	if err := ctx.ShouldBindJSON(&emailData); err != nil {
		c.logger.Fatalf("Cannot send email due to error: %v", err)

		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if len(emailData.To) == 0 {
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Email to cannot be empty")
		return
	}

	err := c.sesService.SendEmail(&services.EmailParams{
		From:    c.env.AdminEmail,
		To:      emailData.To,
		Subject: emailData.Subject,
		Body:    emailData.Body,
	})

	if err != nil {
		c.logger.Fatalf("Cannot send email due to error: %v", err)

		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    fmt.Sprintf("Email sent to %s successfully", emailData.To),
	})
}
