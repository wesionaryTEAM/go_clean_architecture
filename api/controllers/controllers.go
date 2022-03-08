package controllers

import (
	"clean-architecture/api_errors"
	"clean-architecture/lib"
	"errors"
	"net/http"
	"reflect"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserController),
)

func handleValidationError(logger lib.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func handleErrorWithStatus(logger lib.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

// list static errors to filter
var exceptStaticError = []error{
	gorm.ErrRecordNotFound,
	api_errors.ErrInvalidUUID,
}

// list dyanmic errors to filter
var exceptDynamicError = []error{}

// list SQL errors to filter
var exceptSQLError = []uint16{
	1062, // duplicate entry
}

var sqlError *mysql.MySQLError

func handleError(logger lib.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	for _, e := range exceptStaticError {
		if errors.Is(err, e) {
			return
		}
	}

	for _, e := range exceptDynamicError {
		if reflect.TypeOf(e) == reflect.TypeOf(err) {
			return
		}
	}

	if errors.As(err, &sqlError) {
		for _, code := range exceptSQLError {
			if code == sqlError.Number {
				return
			}
		}
	}

	sentry.CaptureException(err)
}
