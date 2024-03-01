package pkg

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
	"clean-architecture/pkg/services"

	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	fx.Options(
		fx.Provide(
			framework.NewEnv,
			framework.GetLogger,
		),
	),
	services.Module,
	infrastructure.Module,
	middlewares.Module,
)
