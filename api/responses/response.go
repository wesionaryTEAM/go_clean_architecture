package responses

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"prototype2/errors"
)

func HandleError(c *gin.Context, err error) {
	// Send error to sentry
	sentry.CaptureMessage(err.Error())

	// Get ErrorType
	errorType := errors.GetErrorType(err)

	// Get Status Code of the ErrorType
	status := errors.GetStatusCode(errorType)

	// Check if there's additional context to the error
	errorContext := errors.GetErrorContext(err)
	if errorContext != nil {
		c.JSON(status, gin.H{
			"error":   err.Error(),
			"context": errorContext,
		})
		return
	}

	// No error context to the error
	c.JSON(status, gin.H{"error": err.Error()})
}
