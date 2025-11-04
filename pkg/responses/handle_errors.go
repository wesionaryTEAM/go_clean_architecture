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

// HandleValidationError returns standardized validation error
func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)
	Error(c, errorz.ErrBadRequest.WithDetail("VALIDATION_ERROR", err.Error()))
}

// HandleErrorWithStatus wraps arbitrary status into API error
func HandleErrorWithStatus(logger framework.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	api := errorz.New("CUSTOM_ERROR", statusCode, http.StatusText(statusCode)).Wrap(err)
	Error(c, api)
}

// HandleError central error translator
func HandleError(logger framework.Logger, c *gin.Context, err error) {
	if err == nil {
		Error(c, errorz.ErrInternal)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Error(c, errorz.ErrRecordNotFound.Wrap(err))
		return
	}
	api := errorz.From(err)
	if api.Code == errorz.ErrInternal.Code && api.Cause != nil {
		api.Details = nil
	}
	Error(c, api)
	if api.StatusCode >= http.StatusInternalServerError {
		utils.CurrentSentryService.CaptureException(err)
	}
}
