package errors

import (
	"net/http"
	// "errors"

	"github.com/gin-gonic/gin"
)

// type ErrResponse struct {
// 	Message string `json:"message"`
// 	Status  int    `json:"status"`
// }

var ErrHTTPStatusMap = map[string]int{
	ErrNotFound.Error():         http.StatusNotFound,
	ErrInvalidSlug.Error():      http.StatusBadRequest,
	ErrExists.Error():           http.StatusConflict,
	ErrDatabase.Error():         http.StatusInternalServerError,
	ErrUnauthorized.Error():     http.StatusUnauthorized,
	ErrForbidden.Error():        http.StatusForbidden,
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     http.StatusBadRequest,
	ErrUserExists.Error():       http.StatusConflict,
}

func Wrap(c *gin.Context, err error) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	// like a default case
	if code == 0 {
		code = http.StatusInternalServerError
		c.JSON(code, gin.H{"error": "Internal Server Error"})
		return
	}

	// errResponse := ErrResponse{
	// 	Message: msg,
	// 	Status:  code,
	// }
	// log.WithFields(log.Fields{
	// 	"message": msg,
	// 	"code":    code,
	// }).Error("Error occurred")

	c.JSON(code, gin.H{"error": msg})
}
