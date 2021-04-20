package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger sets up logger
func NewLogger(env Env) Logger {

	config := zap.NewDevelopmentConfig()

	if env.Environment == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Level.SetLevel(zap.PanicLevel)
	}

	logger, _ := config.Build()

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
