package services

import (
	"clean-architecture/domain/constants"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"

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
	env    *framework.Env
	logger framework.Logger
}

func NewCognitoAuthService(
	client *cognitoidentityprovider.Client,
	env *framework.Env,
	logger framework.Logger,
) CognitoAuthService {

	issuer = "https://cognito-idp." + env.AWSRegion + ".amazonaws.com/" + env.UserPoolID
	jwkURL = issuer + "/.well-known/jwks.json"

	keySet, _ = jwk.Fetch(context.Background(), jwkURL)

	return CognitoAuthService{
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

func (cg *CognitoAuthService) CreateUser(email, password, role string) (string, error) {
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
			return "", awsErr
		}
		return "", err
	}

	_, err = cg.client.AdminSetUserPassword(context.Background(), &cognitoidentityprovider.AdminSetUserPasswordInput{
		Username:   &email,
		Password:   &password,
		Permanent:  true,
		UserPoolId: &cg.env.UserPoolID,
	})
	if err != nil {
		_, delErr := cg.client.AdminDeleteUser(context.Background(), &cognitoidentityprovider.AdminDeleteUserInput{
			Username:   &email,
			UserPoolId: &cg.env.UserPoolID,
		})
		awsErr := utils.MapAWSError(cg.logger, delErr)
		if awsErr != nil {
			return "", awsErr
		}
		return "", utils.MapAWSError(cg.logger, err)
	}
	return "nil", nil
}

func (cg *CognitoAuthService) setCustomClaimToOneUser(user string, c map[string]string) error {
	var claim = make([]types.AttributeType, 0, 5)
	var create = make([]types.SchemaAttributeType, 0, 5)

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

func (cg *CognitoAuthService) GetUserByUsername(username string) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	user, err := cg.client.AdminGetUser(context.Background(), &cognitoidentityprovider.AdminGetUserInput{
		Username:   &username,
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

func (cg *CognitoAuthService) CreateAdminUser(email, password string, isPermanent bool) (string, error) {
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
			return "", awsErr
		}
		return "", err
	}
	_, err = cg.client.AdminSetUserPassword(context.Background(), &cognitoidentityprovider.AdminSetUserPasswordInput{
		Username:   &email,
		Password:   &password,
		Permanent:  true,
		UserPoolId: &cg.env.UserPoolID,
	})

	if err != nil {
		_, err = cg.client.AdminDeleteUser(context.Background(), &cognitoidentityprovider.AdminDeleteUserInput{Username: &email, UserPoolId: &cg.env.UserPoolID})
		awsErr := utils.MapAWSError(cg.logger, err)
		if awsErr != nil {
			return "", awsErr
		}
		return "", err
	}

	err = cg.setCustomClaimToOneUser(email, map[string]string{
		"role":            string(constants.RoleIsAdmin),
		"change-password": strconv.FormatBool(!isPermanent),
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return "", awsErr
		}
		return "", err
	}
	// fetching created admin user
	adminUser, err := cg.GetUserByEmail(email)
	if err != nil {
		cg.logger.Error(err)
		return "", err
	}

	var cognitoUUID string
	// Access the user attributes tp find sub i.e cognito internal uuid
	for _, attr := range adminUser.UserAttributes {
		if *attr.Name == "sub" {
			cognitoUUID = *attr.Value
			break
		}
	}

	return cognitoUUID, nil
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

// UpdateUserAttribute updates user attribute from user's access token
func (cg *CognitoAuthService) UpdateUserAttribute(username *string, attr []types.AttributeType) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	op, err := cg.client.AdminUpdateUserAttributes(context.Background(),
		&cognitoidentityprovider.AdminUpdateUserAttributesInput{
			UserPoolId:     &cg.env.UserPoolID,
			Username:       username,
			UserAttributes: attr,
		},
	)
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return nil, awsErr
		}
		return nil, err
	}
	return op, nil
}

// UpdateEmailAddress update email address of the user by checking
// if the proper password is provided or not
func (cg *CognitoAuthService) UpdateEmailAddress(uid, token, password, email *string) error {
	_, err := cg.client.ChangePassword(context.Background(), &cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      token,
		PreviousPassword: password,
		ProposedPassword: password,
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}

	_, err = cg.UpdateUserAttribute(uid, []types.AttributeType{
		{
			Name:  aws.String("email"),
			Value: email,
		},
	})
	return err
}

// SetUserPassword sets cognito users password from admin
func (cg *CognitoAuthService) SetUserPassword(email, password string) error {
	_, err := cg.client.AdminSetUserPassword(context.Background(), &cognitoidentityprovider.AdminSetUserPasswordInput{
		Password:   &password,
		Username:   &email,
		Permanent:  true,
		UserPoolId: &cg.env.UserPoolID,
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	return nil
}

func (cg *CognitoAuthService) DeleteUserAsAdmin(username string) error {
	_, err := cg.client.AdminDeleteUser(context.Background(),
		&cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: &cg.env.UserPoolID,
			Username:   &username,
		},
	)
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	return nil
}

func (cg *CognitoAuthService) UpdateUserRole(email, newRole string) error {
	err := cg.setCustomClaimToOneUser(email, map[string]string{
		"role": newRole,
	})
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	return nil
}

func (cg *CognitoAuthService) DisableUser(username string) error {
	_, err := cg.client.AdminDisableUser(context.Background(),
		&cognitoidentityprovider.AdminDisableUserInput{
			UserPoolId: &cg.env.UserPoolID,
			Username:   &username,
		},
	)
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	return nil
}

func (cg *CognitoAuthService) EnableUser(username string) error {
	_, err := cg.client.AdminEnableUser(context.Background(),
		&cognitoidentityprovider.AdminEnableUserInput{
			UserPoolId: &cg.env.UserPoolID,
			Username:   &username,
		},
	)
	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return awsErr
		}
		return err
	}
	return nil
}

func (cg *CognitoAuthService) AdminLogin(email string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	out, err := cg.client.AdminInitiateAuth(context.Background(), &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   &cg.env.ClientID,
		UserPoolId: &cg.env.UserPoolID,
		AuthFlow:   types.AuthFlowTypeCustomAuth,
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": "",
		},
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}
