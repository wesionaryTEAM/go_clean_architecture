package services

import (
	"clean-architecture/pkg/framework"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var jwkURL = ""
var issuer = ""
var keySet jwk.Set = jwk.NewSet()

type CognitoAuthService struct {
	client *cognitoidentityprovider.Client
	env    *framework.Env
	logger framework.Logger
}

func NewCognitoAuthService(
	client *cognitoidentityprovider.Client,
	env *framework.Env,
	logger framework.Logger,
) *CognitoAuthService {

	issuer = "https://cognito-idp." + env.AWSRegion + ".amazonaws.com/" + env.AWSCognitoUserPoolID
	jwkURL = issuer + "/.well-known/jwks.json"

	keySet, _ = jwk.Fetch(context.Background(), jwkURL)

	return &CognitoAuthService{
		client: client,
		env:    env,
		logger: logger,
	}
}

func (cg *CognitoAuthService) VerifyToken(tokenString string) (jwt.Token, error) {
	parsedToken, err := jwt.Parse(
		[]byte(tokenString),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.WithIssuer(issuer),
	)

	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
