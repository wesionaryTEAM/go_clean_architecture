package seeds

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"clean-architecture/services"

	"github.com/gin-gonic/gin"
)

// AdminSeed Admin seeding
type AdminSeed struct {
	logger          lib.Logger
	firebaseService services.FirebaseService
	env             *lib.Env
}

// NewAdminSeed Admin seeding
func NewAdminSeed(
	logger lib.Logger,
	firebaseService services.FirebaseService,
	env *lib.Env,
) AdminSeed {
	return AdminSeed{
		logger:          logger,
		firebaseService: firebaseService,
		env:             env,
	}
}

// Run admin seeder
func (as AdminSeed) Setup() {
	// Create email manually in firebase
	email := as.env.AdminEmail
	password := as.env.AdminPassword

	as.logger.Info("🌱 seeding  admin data...")

	if email == "" || password == "" {
		as.logger.Error("Got empty admin email and password from environment variables. Admin seed not executed.")
		return
	}

	_, err := as.firebaseService.RetrieveUserByEmail(email)

	if err != nil {
		adminClaim := gin.H{
			constants.RoleIsAdmin: true,
		}
		_, err := as.firebaseService.CreateUserWithClaims(email, password, adminClaim)
		if err != nil {
			as.logger.Error("Firebase Admin user can't be created: ", err.Error())
			return
		}

		as.logger.Info("Firebase Admin User Created, email: ", email, " password: ", password)
		return
	}

	as.logger.Info("Admin already exist")
}
