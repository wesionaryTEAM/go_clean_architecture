package middleware

import (
	"strings"

	"prototype2/api/responses"
	"prototype2/api/service"
	"prototype2/errors"

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
		err := errors.InternalError.New("no token found in header")
		responses.HandleError(c, err)
		c.Abort()
		return
	}

	token, err := m.authProvider.VerifyToken(idToken)
	if err != nil {
		err := errors.Unauthorized.New("token is not valid")
		responses.HandleError(c, err)
		c.Abort()
		return
	}

	if token.Claims[m.claims] != true {
		err := errors.Forbidden.New("permission denied")
		responses.HandleError(c, err)
		c.Abort()
		return
	}

	c.Next()
}
