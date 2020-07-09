package utils

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
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
