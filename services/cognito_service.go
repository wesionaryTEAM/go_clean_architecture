package services

import (
	"clean-architecture/lib"
	"clean-architecture/utils"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var jwkURL = ""
var issuer = ""
var keySet jwk.Set = jwk.NewSet()

type CognitoAuthService struct {
	client *cognitoidentityprovider.Client
	env    *lib.Env
	logger lib.Logger
}

func NewCognitoAuthService(
	client *cognitoidentityprovider.Client,
	env *lib.Env,
	logger lib.Logger) *CognitoAuthService {

	issuer = "https://cognito-idp." + env.AWSRegion + ".amazonaws.com/" + env.UserPoolID
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

func (cg *CognitoAuthService) CreateUser(email string, password string) (*cognitoidentityprovider.SignUpOutput, error) {
	cognitoUser, err := cg.client.SignUp(context.Background(), &cognitoidentityprovider.SignUpInput{
		ClientId: &cg.env.ClientID,
		Username: &email,
		Password: &password,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	})

	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return nil, awsErr
		}
		return nil, err
	}

	return cognitoUser, nil
}

func (cg *CognitoAuthService) SetCustomClaimToOneUser(user string, c map[string]string) error {
	var claim []types.AttributeType
	var create []types.SchemaAttributeType
	developerOnly := false
	mutable := true
	required := false

	for key, val := range c {

		attribute := types.AttributeType{
			Name:  aws.String("custom:" + key),
			Value: aws.String(val),
		}

		schemaAttribute := types.SchemaAttributeType{
			AttributeDataType:      "String",
			DeveloperOnlyAttribute: &developerOnly,
			Mutable:                &mutable,
			Name:                   aws.String(key),
			Required:               &required,
		}
		claim = append(claim, attribute)
		create = append(create, schemaAttribute)
	}

	_, _ = cg.client.AddCustomAttributes(context.Background(), &cognitoidentityprovider.AddCustomAttributesInput{
		CustomAttributes: create,
		UserPoolId:       &cg.env.UserPoolID,
	})

	_, err := cg.client.AdminUpdateUserAttributes(context.Background(), &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: claim,
		UserPoolId:     &cg.env.UserPoolID,
		Username:       &user,
	})

	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}

	return nil
}

func (cg *CognitoAuthService) GetUserByEmail(email string) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	user, err := cg.client.AdminGetUser(context.Background(), &cognitoidentityprovider.AdminGetUserInput{
		Username:   &email,
		UserPoolId: &cg.env.UserPoolID,
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return nil, awsErr
		}
		return nil, err
	}

	return user, nil
}

func (cg *CognitoAuthService) CreateAdminUser(email, password string, isPermanent bool) error {
	_, err := cg.client.AdminCreateUser(context.Background(), &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:    &cg.env.UserPoolID,
		Username:      &email,
		MessageAction: types.MessageActionTypeSuppress,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
		ValidationData: []types.AttributeType{},
	})

	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	_, err = cg.client.AdminSetUserPassword(context.Background(), &cognitoidentityprovider.AdminSetUserPasswordInput{

		Username:   &email,
		Password:   &password,
		Permanent:  true,
		UserPoolId: &cg.env.UserPoolID,
	})

	if err != nil {
		_, err := cg.client.AdminDeleteUser(context.Background(), &cognitoidentityprovider.AdminDeleteUserInput{Username: &email, UserPoolId: &cg.env.UserPoolID})
		awsErr := utils.MapAWSError(cg.logger, err)
		if awsErr != nil {
			return awsErr
		}
		return err
	}

	err = cg.SetCustomClaimToOneUser(email, map[string]string{
		"role":            "admin",
		"change-password": strconv.FormatBool(!isPermanent),
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}

	return nil
}

func (cg *CognitoAuthService) DeleteCognitoUser(token *string) error {
	_, err := cg.client.DeleteUser(context.Background(), &cognitoidentityprovider.DeleteUserInput{
		AccessToken: token,
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}

	return nil
}
