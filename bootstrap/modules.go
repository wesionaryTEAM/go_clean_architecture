package bootstrap

import (
	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/api/routes"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
	"clean-architecture/repository"
	"clean-architecture/seeds"
	"clean-architecture/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	infrastructure.Module,
	middlewares.Module,
	lib.Module,
	seeds.Module,
)
