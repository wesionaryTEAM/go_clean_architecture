package bootstrap

import (
	"context"

	"go.uber.org/fx"
	"prototype2/api/handlers"
	"prototype2/api/middlewares"
	"prototype2/api/repository"
	"prototype2/api/routes"
	"prototype2/api/services"
	"prototype2/cli"
	"prototype2/lib"
	"prototype2/utils"
)

var Module = fx.Options(
	handlers.Module,
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
				logger.Zap.Info("------- 🤖 hatsu 🤖 (CLI) -------")
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
			logger.Zap.Info("------------------------")
			logger.Zap.Info("------- hatsu 📺 -------")
			logger.Zap.Info("------------------------")

			logger.Zap.Info("Migrating DB schema...")

			//migrations.Migrate()

			go func() {
				middlewares.Setup()
				routes.Setup()
				logger.Zap.Info("🌱 seeding data...🌱 🌱 🌱 ")
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()

			return nil
		},
		OnStop: appStop,
	})
}
