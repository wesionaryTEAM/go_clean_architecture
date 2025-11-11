package responses

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

// HandleValidationError returns standardized validation error
func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)
	var fieldErrors []map[string]any
	if errs, ok := err.(validation.Errors); ok {
		for field, ferr := range errs {
			if ferr == nil {
				continue
			}
			fieldErrors = append(fieldErrors, map[string]any{
				"field":   field,
				"message": ferr.Error(),
			})
		}
	} else {
		fieldErrors = append(fieldErrors, map[string]any{
			"field":      "",
			"error_type": "invalid",
			"message":    err.Error(),
		})
	}

	Error(c, errorz.ErrBadRequest.WithDetail("validation_errors", fieldErrors))
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
