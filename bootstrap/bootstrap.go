package bootstrap

import (
	"context"

	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/api/repository"
	"clean-architecture/api/routes"
	"clean-architecture/api/services"
	"clean-architecture/cli"
	"clean-architecture/lib"
	"clean-architecture/utils"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	lib.Module,
	middlewares.Module,
	cli.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.Router,
	routes routes.Routes,
	env lib.Env,
	middlewares middlewares.Middlewares,
	logger lib.Logger,
	cliApp cli.Application,
	database lib.Database,

) {

	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")

		conn, _ := database.DB.DB()
		conn.Close()
		return nil
	}

	if utils.IsCli() {
		lifecycle.Append(fx.Hook{
			OnStart: func(context.Context) error {
				logger.Zap.Info("Starting hatsu cli Application")
				logger.Zap.Info("------- 🤖 clean-architecture 🤖 (CLI) -------")
				go cliApp.Start()
				return nil
			},
			OnStop: appStop,
		})

		return
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-------------------------------------")
			logger.Zap.Info("------- clean-architecture 📺 -------")
			logger.Zap.Info("-------------------------------------")

			go func() {
				middlewares.Setup()
				routes.Setup()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(":" + env.ServerPort)
				}
			}()

			return nil
		},
		OnStop: appStop,
	})
}
