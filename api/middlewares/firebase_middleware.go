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
func (m FirebaseAuthMiddleware) HandleAuthWithRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := m.getTokenFromHeader(ctx)
		if err != nil {
			responses.ErrorJSON(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		// verify user role
		if len(roles) > 0 {
			if ok := m.checkIsRoleSatisfied(roles, token); !ok {
				responses.ErrorJSON(ctx, http.StatusForbidden, "auth-not-authorized-user")
				ctx.Abort()
				return
			}
		}

		ctx.Set(constants.Claims, token.Claims)
		ctx.Set(constants.UID, token.UID)

		ctx.Next()
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

func (m FirebaseAuthMiddleware) checkIsRoleSatisfied(roles []string, token *auth.Token) bool {
	for _, val := range roles {
		if token.Claims[val] == true {
			return true
		}
	}
	return false
}
