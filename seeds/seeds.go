package seeds

import "go.uber.org/fx"

// Module exports seed module
var Module = fx.Options(
	fx.Provide(NewAdminSeed),
	fx.Provide(NewSeeds),
)

// Seed db seed
type Seed interface {
	// name is used to identify the seed
	Name() string
	Setup()
}

// Seeds listing of seeds
type Seeds []Seed

// NewSeeds creates new seeds
func NewSeeds(adminSeed AdminSeed) Seeds {
	return Seeds{
		adminSeed,
	}
}
