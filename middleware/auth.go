package middleware

import (
	"net/http"
	"strings"

	"prototype2/service"

	"github.com/gin-gonic/gin"
)

type middlewareAuth struct {
	authProvider service.FirebaseService
}

// NewMiddlewareAuth : get injected firebase service
func NewMiddlewareAuth(s service.FirebaseService) gin.HandlerFunc {
	return (&middlewareAuth{
		authProvider: s,
	}).UserAuth
}

// UserAuth : to verify all authorized operations
func (m *middlewareAuth) UserAuth(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")

	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No token found in header",
		})
		c.Abort()
		return
	}

	_, err := m.authProvider.VerifyToken(idToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "token is not valid",
		})
		c.Abort()
		return

	}

	c.Next()
}
