package seeds

import (
	"clean-architecture/pkg/framework"
)

// AdminSeed Admin seeding
type AdminSeed struct {
	logger framework.Logger
	env    *framework.Env
}

// NewAdminSeed Admin seeding
func NewAdminSeed(
	logger framework.Logger,
	env *framework.Env,
) AdminSeed {
	return AdminSeed{
		logger: logger,
		env:    env,
	}
}

// Run admin seeder
func (as AdminSeed) Setup() {
	// Create email manually in firebase

}
