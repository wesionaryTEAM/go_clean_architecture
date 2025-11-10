package responses

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// HandleValidationError returns standardized validation error
func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)

	var fieldErrors []map[string]any
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			fieldErrors = append(fieldErrors, map[string]any{
				"field":      fe.Field(),
				"error_type": fe.Tag(),
				"message":    fe.Error(),
			})
		}
	} else {
		fieldErrors = append(fieldErrors, map[string]any{
			"field":      "",
			"error_type": "invalid",
			"message":    err.Error(),
		})
	}

	// BAD_REQUEST with our details
	responsesError := errorz.ErrBadRequest.WithDetail("validation_errors", fieldErrors)
	Error(c, responsesError)
}

// HandleErrorWithStatus wraps arbitrary status into API error
func HandleErrorWithStatus(logger framework.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	api := errorz.New(errorz.CodeCustomError, statusCode, http.StatusText(statusCode)).Wrap(err)
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
