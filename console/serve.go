package console

import (
	"clean-architecture/api/middlewares"
	"clean-architecture/api/routes"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env *lib.Env,
		router infrastructure.Router,
		route *routes.Routes,
		logger lib.Logger,
	) {
		logger.Info(`+-----------------------+`)
		logger.Info(`| GO CLEAN ARCHITECTURE |`)
		logger.Info(`+-----------------------+`)
		middleware.Setup()
		route.Setup()
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
		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Run()
		} else {
			_ = router.Run(":" + env.ServerPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
