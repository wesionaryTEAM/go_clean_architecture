package middlewares

import (
	"clean-architecture/api/responses"
	"clean-architecture/constants"
	"clean-architecture/services"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// FirebaseAuthMiddleware structure
type FirebaseAuthMiddleware struct {
	service services.FirebaseService
}

// NewFirebaseAuthMiddleware creates new firebase authentication
func NewFirebaseAuthMiddleware(service services.FirebaseService) FirebaseAuthMiddleware {
	return FirebaseAuthMiddleware{
		service: service,
	}
}

// HandleAuthWithRole handles multiple roles
func (m FirebaseAuthMiddleware) HandleAuthWithRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := m.getTokenFromHeader(c)
		if err != nil {
			responses.ErrorJSON(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if role != "" && token.Claims[role] == nil {
			responses.ErrorJSON(c, http.StatusUnauthorized, "auth-not-authorized-user")
			c.Abort()
			return
		}

		c.Set(constants.Claims, token.Claims)
		c.Set(constants.UID, token.UID)

		c.Next()
	}
}

// getTokenFromHeader gets token from header
func (m FirebaseAuthMiddleware) getTokenFromHeader(c *gin.Context) (*auth.Token, error) {
	header := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := m.service.VerifyToken(idToken)
	if err != nil {
		return nil, err
	}

	// set email to the sentry message
	email := token.Claims["email"]
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{Email: email.(string)})
		})
	}

	return token, nil
}
