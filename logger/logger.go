package logger

import (
	"time"

	"go.ssnk.in/utils/logger/constants"
	"go.ssnk.in/utils/logger/internal"
	"go.ssnk.in/utils/logger/ports"
)

type Logger struct {
	ports.Logger

	provider   constants.Provider
	level      constants.Level
	format     constants.Format
	trace      bool
	rotateOpts RotateOptions
}

type RotateOptions struct {
	format      constants.Format
	duration    time.Duration
	maxFileSize int64
}

func New(opts ...func(*Logger)) Logger {
	l := &Logger{
		Logger: nil,
	}

	for _, opt := range opts {
		opt(l)
	}

	var logger ports.Logger

	switch l.provider {
	case constants.Logrus:
	case constants.Zap:
	case constants.Zerolog:
	default:
		logger = internal.NewSlogLogger(l.provider.String(), l.level.String(), l.trace)
	}

	l.Logger = logger

	return *l
}

func SetProvider(provider constants.Provider) func(*Logger) {
	return func(logger *Logger) {
		if len(provider) == 0 {
			provider = constants.Slog
		}

		logger.provider = provider
	}
}

func SetLevel(level constants.Level) func(*Logger) {
	return func(logger *Logger) {
		if len(level) == 0 {
			level = constants.Debug
		}

		logger.level = level
	}
}

func SetFormat(format constants.Format) func(*Logger) {
	return func(logger *Logger) {
		if len(format) == 0 {
			format = constants.Json
		}

		logger.format = format
	}
}

func WithTracing() func(*Logger) {
	return func(logger *Logger) {
		logger.trace = true
	}
}

func WithRotateOptions(duration, maxFileSize int64, format string) func(*Logger) {
	return func(logger *Logger) {
		if duration == 0 {
			duration = int64(24 * time.Hour)
		}

		if maxFileSize == 0 {
			maxFileSize = 8 * 1024 * 1024
		}

		if len(format) == 0 {
			format = string(constants.Proto)
		}

		logger.rotateOpts.duration = time.Duration(duration)
		logger.rotateOpts.maxFileSize = maxFileSize
		logger.rotateOpts.format = constants.Format(format)
	}
}
