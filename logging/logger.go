package logging

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger

	level zap.AtomicLevel
)

// loggerContextKey is a zero-allocation, private key used to store the logger in a context.
type loggerContextKey struct{}

const (
	// DefaultLevel is the default log level used if none is provided
	defaultLevel = zapcore.InfoLevel
)

// Initialize initializes a global zap logger.
// it returns an already initialized logger if called multiple times and updates the log level
func Initialize(l zapcore.Level) *zap.Logger {
	if globalLogger != nil {
		level.SetLevel(l)
		return globalLogger
	}

	level = zap.NewAtomicLevelAt(l)

	logConfig := zap.NewProductionConfig()
	logConfig.Level = level
	logConfig.EncoderConfig.MessageKey = "message"
	logConfig.EncoderConfig.TimeKey = "@timestamp"
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	globalLogger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}

	globalLogger.Info("logger initialized")

	return globalLogger
}

// InitializeFromEnv reads the log level from the SERVICE_LOG_LEVEL environment variable and initializes a logger
func InitializeFromEnv() *zap.Logger {
	level := strings.ToUpper(os.Getenv("SERVICE_LOG_LEVEL"))

	switch level {
	case "DEBUG", "D":
		return Initialize(zap.DebugLevel)
	case "INFO", "I":
		return Initialize(zap.InfoLevel)
	case "WARN", "W":
		return Initialize(zap.WarnLevel)
	case "ERROR", "E":
		return Initialize(zap.ErrorLevel)
	}
	return Initialize(defaultLevel)
}

// logger returns a logger from the context, sets a name and adds any passed fields into it.
func logger(ctx context.Context, name string, fields ...zap.Field) *zap.Logger {
	//Check if logger is initialized
	if globalLogger == nil {
		InitializeFromEnv()
	}

	if ctx == nil {
		return globalLogger.Named(name).With(fields...)
	}

	if ctxLogger, ok := ctx.Value(loggerContextKey{}).(*zap.Logger); ok {
		return ctxLogger.Named(name).With(fields...)
	}
	return globalLogger.Named(name).With(fields...)
}

func Logger(ctx context.Context, name string, fields ...zap.Field) (context.Context, *zap.Logger) {
	l := logger(ctx, name, fields...)
	return context.WithValue(ctx, loggerContextKey{}, l), l
}

func WrapContext(ctx context.Context, name string, fields ...zap.Field) context.Context {
	c, _ := Logger(ctx, name, fields...)
	return c
}
