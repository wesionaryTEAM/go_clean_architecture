package responses

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"errors"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

// HandleError is the central error translator.
func HandleError(logger framework.Logger, c *gin.Context, err error) {
	var apiErr *errorz.APIError
	if errors.As(err, &apiErr) {
		Error(c, apiErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Error(c, errorz.ErrRecordNotFound)
		return
	}
	logger.Error(err)
	utils.CurrentSentryService.CaptureException(err)
	Error(c, errorz.ErrInternal)
}

// HandleErrorWithParams writes an APIError with the given params attached.
func HandleErrorWithParams(logger framework.Logger, c *gin.Context, apiErr *errorz.APIError, params map[string]any) {
	Error(c, apiErr.WithParams(params))
}

// HandleErrorWithStatus wraps an arbitrary status into an API error.
func HandleErrorWithStatus(logger framework.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	severity := errorz.SeverityWarn
	if statusCode >= http.StatusInternalServerError {
		severity = errorz.SeverityError
		utils.CurrentSentryService.CaptureException(err)
	}
	Error(c, errorz.New(errorz.CodeCustomError, statusCode, severity).Wrap(err))
}

// HandleValidationError returns a standardized validation error.
func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	// An *APIError is a deliberate, already-classified error (e.g. a 400-class
	// client error); pass it through without logging it as an internal error.
	var apiErr *errorz.APIError
	if errors.As(err, &apiErr) {
		Error(c, apiErr)
		return
	}

	logger.Error(err)

	var verrs validation.Errors
	if errors.As(err, &verrs) {
		fieldErrors := make([]map[string]any, 0, len(verrs))
		for field, ferr := range verrs {
			if ferr == nil {
				continue
			}
			fieldErrors = append(fieldErrors, map[string]any{
				"field":      field,
				"error_type": "validation",
				"message":    ferr.Error(),
			})
		}
		// Sort by field so the response order is deterministic (map iteration
		// order is random in Go).
		sort.Slice(fieldErrors, func(i, j int) bool {
			return fieldErrors[i]["field"].(string) < fieldErrors[j]["field"].(string)
		})
		Error(c, errorz.ErrBadRequest.WithParam("validation_errors", fieldErrors))
		return
	}

	HandleError(logger, c, err)
}
