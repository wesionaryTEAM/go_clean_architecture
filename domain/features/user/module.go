package user

import (
	"clean-architecture/domain/domainif"

	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Options(
		fx.Provide(
			NewRepository,
			NewController,
			NewRoute,
			fx.Annotate(NewService, fx.As(new(domainif.UserService))),
		),
		fx.Invoke(RegisterRoute),
	))
