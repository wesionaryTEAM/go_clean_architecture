package middlewares

import (
	"clean-architecture/pkg/api_errors"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/services"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwt"
)

type CognitoAuthMiddleware struct {
	service services.CognitoAuthService
}

func NewCognitoAuthMiddleware(service services.CognitoAuthService) CognitoAuthMiddleware {
	return CognitoAuthMiddleware{
		service: service,
	}
}

func (am CognitoAuthMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := am.addClaimsToContext(ctx); err != nil {
			responses.ErrorJSON(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
	}
}

func (am CognitoAuthMiddleware) getTokenFromHeader(ctx *gin.Context) (jwt.Token, error) {
	header := ctx.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := am.service.VerifyToken(idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (am CognitoAuthMiddleware) addClaimsToContext(ctx *gin.Context) error {
	token, err := am.getTokenFromHeader(ctx)
	if err != nil {
		return api_errors.ErrUnauthorizedAccess
	}
	claims := token.PrivateClaims()

	ctx.Set(framework.Claims, claims)
	username := claims["cognito:username"]
	ctx.Set(framework.UID, username)

	role, ok := claims["custom:role"]
	if ok {
		ctx.Set(framework.Role, role)
	}
	ctx.Set(framework.CognitoPass, true)
	return nil
}
