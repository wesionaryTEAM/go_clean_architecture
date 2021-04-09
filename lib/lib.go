package lib

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewRouter),
	fx.Provide(NewEnv),
	fx.Provide(NewDatabase),
)
