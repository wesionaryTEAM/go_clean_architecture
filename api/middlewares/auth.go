package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"clean-architecture/api/responses"
	"clean-architecture/api/services"

	"github.com/gin-gonic/gin"
)

type middlewareAuth struct {
	authProvider services.FirebaseService
	claims       string
}

// NewMiddlewareAuth : get injected firebase service and claims
func NewMiddlewareAuth(s services.FirebaseService, c string) gin.HandlerFunc {
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
		err := errors.New("no token found in header")
		responses.ErrorJSON(c, http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	token, err := m.authProvider.VerifyToken(idToken)
	if err != nil {
		err := errors.New("token is not valid")
		responses.ErrorJSON(c, http.StatusUnauthorized, err)
		c.Abort()
		return
	}

	if token.Claims[m.claims] != true {
		err := errors.New("permission denied")
		responses.ErrorJSON(c, http.StatusForbidden, err)
		c.Abort()
		return
	}

	c.Next()
}
