package lib

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gormlogger "gorm.io/gorm/logger"
)

var globalLog *Logger

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

type GormLogger struct {
	*Logger
	gormlogger.Config
}

// GetLogger gets the global instance of the logger
func GetLogger() Logger {
	if globalLog != nil {
		return *globalLog
	}
	globalLog := newLogger()
	return *globalLog
}

// NewLogger sets up logger
func newLogger() *Logger {

	env := os.Getenv("ENVIRONMENT")
	config := zap.NewDevelopmentConfig()

	if env == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Level.SetLevel(zap.PanicLevel)
	}

	logger, _ := config.Build()

	globalLog := logger.Sugar()

	return &Logger{
		SugaredLogger: globalLog,
	}

}

// GetGormLogger build gorm logger from zap logger
func (l *Logger) GetGormLogger() gormlogger.Interface {
	return &GormLogger{
		Logger: l,
		Config: gormlogger.Config{
			LogLevel: gormlogger.Info,
		},
	}
}

// ------ GORM logger interface implementation -----

// LogMode set log mode
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Warnf(str, args...)
	}

}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Errorf(str, args...)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.SugaredLogger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.SugaredLogger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}
}

// Printf prints go-fx logs
func (l Logger) Printf(str string, args ...interface{}) {
	l.Infof(str, args)
}
