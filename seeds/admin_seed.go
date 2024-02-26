package seeds

import (
	"clean-architecture/domain/constants"
	"clean-architecture/domain/models"
	"clean-architecture/domain/user"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/services"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type AdminSeed struct {
	logger         framework.Logger
	cognitoService services.CognitoAuthService
	userService    *user.Service
	env            *framework.Env
}

// NewAdminSeed creates admin seed
func NewAdminSeed(
	logger framework.Logger,
	cognitoService services.CognitoAuthService,
	userService *user.Service,
	env *framework.Env,
) AdminSeed {
	return AdminSeed{
		logger:         logger,
		cognitoService: cognitoService,
		userService:    userService,
		env:            env,
	}
}

// Run the admin seed
func (s AdminSeed) Setup() {
	email := s.env.AdminEmail
	password := s.env.AdminPassword

	s.logger.Info("🌱 seeding admin data...")

	if _, err := s.cognitoService.GetUserByUsername(email); err != nil {
		cognitoUUID, err := s.cognitoService.CreateAdminUser(email, password, true)
		if err != nil {
			s.logger.Error("failed to create the admin user in cognito", err.Error())
			return
		}
		s.logger.Info("Successfully created admin user in cognito")

		adminUser := models.User{
			Email:           email,
			CognitoUID:      aws.String(cognitoUUID),
			Role:            constants.RoleIsAdmin,
			IsAdmin:         true,
			IsEmailVerified: true,
			IsActive:        true,
		}
		if err := s.userService.Create(&adminUser); err != nil {
			s.logger.Error(err.Error())
			return
		}
	}
	s.logger.Info("Admin user already exists")
}
