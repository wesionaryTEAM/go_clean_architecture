package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	BadRequest = StatusCode(iota)
	Unauthorized
	Forbidden
	NotFound
	Conflict
	InternalError
	Unavailable
)

type StatusCode uint

type responseError struct {
	statusCode    StatusCode
	originalError error
	contextInfo   errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// Error returns the message of a responseError
func (error responseError) Error() string {
	return error.originalError.Error()
}

// New creates a new responseError
func (code StatusCode) New(msg string) error {
	return responseError{
		statusCode: code,
		originalError: errors.New(msg),
	}
}

// New creates a new responseError with formatted message
func (code StatusCode) Newf(msg string, args ...interface{}) error {
	return responseError{
		statusCode: code,
		originalError: fmt.Errorf(msg, args...),
	}
}

// Wrap creates a new wrapped error
func (code StatusCode) Wrap(err error, msg string) error {
	return code.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (code StatusCode) Wrapf(err error, msg string, args ...interface{}) error {
	return responseError{
		statusCode: code,
		originalError: errors.Wrapf(err, msg, args...),
	}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if responseErr, ok := err.(responseError); ok {
		return responseError{
			statusCode: responseErr.statusCode,
			originalError: responseErr.originalError,
			contextInfo: context,
		}
	}

	return responseError{
		statusCode: InternalError,
		originalError: err,
		contextInfo: context,
	}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if responseErr, ok := err.(responseError); ok && responseErr.contextInfo != emptyContext {

		return map[string]string{
			"field": responseErr.contextInfo.Field,
			"message": responseErr.contextInfo.Message,
		}
	}

	return nil
}

// GetType returns the error type
func GetCode(err error) StatusCode {
	if responseErr, ok := err.(responseError); ok {
		return responseErr.statusCode
	}

	return InternalError
}

// GetType returns the status code for the error type
func GetStatusCode(errorType StatusCode) int {
	switch errorType {
	case BadRequest:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case Conflict:
		return http.StatusConflict
	case InternalError:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func HandleError(c *gin.Context, err error) {
	errorType := GetCode(err)

	status := GetStatusCode(errorType)

	errorContext := GetErrorContext(err)
	if errorContext != nil {
		c.JSON(status, gin.H{
			"error":  err.Error(),
			"context": errorContext,
		})
		return
	}
	c.JSON(status, gin.H{"error": err.Error()})
}
