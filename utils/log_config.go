package utils

import (
	"log"

	"github.com/getsentry/sentry-go"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLumberjackLoging() error {
	err := InitLumberjackLoging()

	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func InitLumberjackLoging() error {
	get := GetEnvWithKey

	MaxSizeInt, _ := ConvertStringToInt(get("LOG_MAX_SIZE"))
	MaxBackupsInt, _ := ConvertStringToInt(get("LOG_MAX_BACKUPS"))
	MaxAgeInt, _ := ConvertStringToInt(get("LOG_MAX_AGE"))

	log.SetOutput(&lumberjack.Logger{
		Filename:   "log/access.log",
		MaxSize:    MaxSizeInt, // megabytes
		MaxBackups: MaxBackupsInt,
		MaxAge:     MaxAgeInt, // days
		Compress:   true,      // disabled by default
	})
	return nil
}
