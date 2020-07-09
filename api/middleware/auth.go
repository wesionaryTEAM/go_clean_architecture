package middleware

import (
	"net/http"
	"strings"

	"prototype2/api/service"

	"github.com/gin-gonic/gin"
)

type middlewareAuth struct {
	authProvider service.FirebaseService
	claims       string
}

// NewMiddlewareAuth : get injected firebase service and claims
func NewMiddlewareAuth(s service.FirebaseService, c string) gin.HandlerFunc {
	return (&middlewareAuth{
		authProvider: s,
		claims:       c,
	}).AuthRequired
}

// AuthRequired : to verify all authorized operations
func (m *middlewareAuth) AuthRequired(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")

	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No token found in header",
		})
		c.Abort()
		return
	}

	token, err := m.authProvider.VerifyToken(idToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "token is not valid",
		})
		c.Abort()
		return

	}

	if token.Claims[m.claims] != true {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Permission denied",
		})
		c.Abort()
		return
	}

	c.Next()
}
