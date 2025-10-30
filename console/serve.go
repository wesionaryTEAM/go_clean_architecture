package console

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() framework.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env *framework.Env,
		router infrastructure.Router,
		logger framework.Logger,
		database infrastructure.Database,
		//seeds seeds.Seeds,

	) {
		logger.Info(`+-----------------------+`)
		logger.Info(`| GO CLEAN ARCHITECTURE |`)
		logger.Info(`+-----------------------+`)

		// Using time zone as specified in env file
		loc, _ := time.LoadLocation(env.Server.TimeZone)
		time.Local = loc

		middleware.Setup()
		//seeds.Setup()

		if env.Server.Environment != "local" && env.Sentry.DSN != "" {
			err := sentry.Init(sentry.ClientOptions{
				Dsn:              env.Sentry.DSN,
				Environment:      env.Server.Environment,
				AttachStacktrace: true,
			})
			if err != nil {
				logger.Error("sentry initialization failed")
				logger.Error(err.Error())
			}
		}
		logger.Info("Running server")
		if env.Server.Port == "" {
			if err := router.Run(); err != nil {
				logger.Fatal(err)
				return
			}
		} else {
			if err := router.Run(":" + env.Server.Port); err != nil {
				logger.Fatal(err)
				return
			}
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
