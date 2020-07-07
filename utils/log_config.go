package utils

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/getsentry/sentry-go"
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
	log.SetOutput(&lumberjack.Logger{
		Filename: "log/access.log",
		MaxSize: 500, // megabytes
		MaxBackups: 3, 
		MaxAge: 28, // days
		Compress: true, // disabled by default
	})
	return nil
}