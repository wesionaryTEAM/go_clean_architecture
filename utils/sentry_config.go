package utils

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SetupSentry() {
	get := GetEnvWithKey
	err := sentry.Init(sentry.ClientOptions{
		Dsn: get("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Sentry setup completed...")
}

//SendSentryMsg -> send  message to sentry
func SendSentryMsg(ctx *gin.Context, msg string) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			hub.CaptureMessage("Error Occured: " + msg)
		})
	}

}
