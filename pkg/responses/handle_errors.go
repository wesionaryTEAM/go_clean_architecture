package responses

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func HandleErrorWithStatus(logger framework.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func HandleError(logger framework.Logger, c *gin.Context, err error) {
	msgForUnhandledError := "An error occurred while processing your request. Please try again later."

	var apiErr *errorz.APIError
	msg := err.Error()
	if ok := errors.As(err, &apiErr); ok {
		if msg == "" {
			msg = apiErr.Message
		}
		c.JSON(apiErr.StatusCode, gin.H{
			"error": msg,
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gorm.ErrRecordNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": msgForUnhandledError,
	})

	utils.CurrentSentryService.CaptureException(err)
}
