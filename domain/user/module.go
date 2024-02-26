package user

import "go.uber.org/fx"

var Module = fx.Module("user",
	fx.Options(
		fx.Provide(
			NewRepository,
			NewService,
			NewController,
			NewRoute,
		),
		fx.Invoke(Migrate, RegisterRoute),
	))
