package utils

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// SendSentryMsg -> send  message to sentry
func SendSentryMsg(ctx *gin.Context, msg string) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			hub.CaptureMessage("Error Occurred: " + msg)
		})
	}

}
