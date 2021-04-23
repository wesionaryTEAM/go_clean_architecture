package bootstrap

import (
	"context"
	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/api/routes"
	"clean-architecture/cli"
	"clean-architecture/infrastructure"
	"clean-architecture/repository"
	"clean-architecture/services"
	"clean-architecture/utils"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	infrastructure.Module,
	middlewares.Module,
	cli.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	env infrastructure.Env,
	middlewares middlewares.Middlewares,
	logger infrastructure.Logger,
	cliApp cli.Application,
	database infrastructure.Database,
	cfg *modcfg.Spec

) {

	var flushTimeout = 2 * time.Second

	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")
		conn, _ := database.DB.DB()
		sentry.Flush(flushTimeout)
		conn.Close()
		return nil
	}

	if utils.IsCli() {
		lifecycle.Append(fx.Hook{
			OnStart: func(context.Context) error {

				if cfg.SentryDSN == "" {
					logger.Info("Sentry disabled")
	
					return nil
				}


				return sentry.Init(sentry.ClientOptions{
					Dsn:         cfg.SentryDSN,
					Environment: cfg.SentryENV,
					Release:     version.AppVersion,
				})
	
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
					handler.Run()
				} else {
					handler.Run(":" + env.ServerPort)
				}
			}()

			return nil
		},
		OnStop: appStop,
	})
}
