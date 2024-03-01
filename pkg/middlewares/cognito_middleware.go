package middlewares

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwt"
)

type CognitoAuthMiddleware struct {
	service *services.CognitoAuthService
}

func NewCognitoAuthMiddleware(service *services.CognitoAuthService) CognitoAuthMiddleware {
	return CognitoAuthMiddleware{
		service: service,
	}
}

func (cm CognitoAuthMiddleware) HandleAuthWithRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := cm.getTokenFromHeader(ctx)
		if err != nil {
			responses.ErrorJSON(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		claims := token.PrivateClaims()

		if len(roles) > 0 {
			if ok := cm.checkIsRoleSatisfied(roles, claims["custom:role"]); !ok {
				responses.ErrorJSON(ctx, http.StatusForbidden, "auth-not-authorized-user")
				ctx.Abort()
				return
			}
		}

		ctx.Set(framework.Claims, claims)
		username := claims["cognito:username"]
		ctx.Set(framework.UID, username)

		ctx.Set(framework.Role, claims["custom:role"])

		ctx.Next()
	}
}

func (cm CognitoAuthMiddleware) getTokenFromHeader(gc *gin.Context) (jwt.Token, error) {
	header := gc.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := cm.service.VerifyToken(idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (cm CognitoAuthMiddleware) checkIsRoleSatisfied(roles []string, role interface{}) bool {
	for _, val := range roles {
		if val == role.(string) {
			return true
		}
	}
	return false
}
