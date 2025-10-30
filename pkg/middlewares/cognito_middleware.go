package middlewares

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/services"
	"fmt"
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
		return errorz.ErrUnauthorizedAccess
	}

	claims := token.PrivateClaims()
	username := claims["cognito:username"]
	authCogUser, err := am.service.GetUserByUsername(fmt.Sprint(username))
	if err != nil {
		return err
	}
	if !authCogUser.Enabled {
		return errorz.ErrUnauthorizedAccess
	}

	ctx.Set(framework.Claims, claims)
	ctx.Set(framework.UID, username)

	role, ok := claims["custom:role"]
	if ok {
		ctx.Set(framework.Role, role)
	}
	ctx.Set(framework.CognitoPass, true)

	return nil
}
