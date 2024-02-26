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
		//If you want to enable auto-migrate add Migrate as shown below
		// fx.Invoke(Migrate, RegisterRoute),

		fx.Invoke(RegisterRoute),
	))
