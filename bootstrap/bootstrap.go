package bootstrap

import (
	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/api/routes"
	"clean-architecture/cmd"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
	"clean-architecture/repository"
	"clean-architecture/services"
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	infrastructure.Module,
	middlewares.Module,
	cmd.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

var flushTimeout = 2 * time.Second

func bootstrap(
	lifecycle fx.Lifecycle,
	middlewares middlewares.Middlewares,
	env lib.Env,
	router infrastructure.Router,
	routes routes.Routes,
	logger lib.Logger,
	database infrastructure.Database,
	rootCmd cmd.RootCommand,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			rootCmd.Run = func(cmd *cobra.Command, args []string) {
				logger.Info(`+-----------------------+`)
				logger.Info(`| GO CLEAN ARCHITECTURE |`)
				logger.Info(`+-----------------------+`)
				middlewares.Setup()
				routes.Setup()
				if env.ServerPort == "" {
					router.Run()
				} else {
					router.Run(":" + env.ServerPort)
				}
				if env.Environment != "local" && env.SentryDSN != "" {
					err := sentry.Init(sentry.ClientOptions{
						Dsn:              env.SentryDSN,
						AttachStacktrace: true,
					})
					if err != nil {
						logger.Error("sentry initialization failed")
						logger.Error(err.Error())
					}
				}
			}
			go rootCmd.Execute()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			sentry.Flush(flushTimeout)
			conn, _ := database.DB.DB()
			conn.Close()
			return nil
		},
	})
}
