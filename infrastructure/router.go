package infrastructure

import (
	"clean-architecture/lib"
	"clean-architecture/utils"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter(
	env *lib.Env,
	logger lib.Logger,
) Router {

	if env.Environment != "local" && env.SentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         env.SentryDSN,
			Environment: `clean-backend-` + env.Environment,
		}); err != nil {
			logger.Infof("Sentry initialization failed: %v\n", err)
		}
	}

	gin.DefaultWriter = logger.GetGinLogger()
	appEnv := env.Environment
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.Default()

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Attach sentry middleware
	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		utils.SendSentryMsg(c, "Error")
		c.JSON(http.StatusOK, gin.H{"data": "clean architecture 📺 API Up and Running"})
	})

	return Router{
		httpRouter,
	}
}
