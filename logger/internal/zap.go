package internal

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction() // Replace with your desired Zap configuration
	return &ZapLogger{
		logger: logger,
	}
}

func (z *ZapLogger) Info(s string, args ...interface{}) {
	z.logger.Sugar().Infof(s, args...)
}

func (z *ZapLogger) Warn(s string, args ...interface{}) {
	z.logger.Sugar().Warnf(s, args...)
}

func (z *ZapLogger) Error(err error, args ...interface{}) {
	z.logger.Sugar().Errorf("%v: %s", err, fmt.Sprint(args...))
}

func (z *ZapLogger) Fatal(err error, args ...interface{}) {
	z.logger.Sugar().Fatalf("%v: %s", err, fmt.Sprint(args...))
}

func (z *ZapLogger) Debug(s string, args ...interface{}) {
	z.logger.Sugar().Debugf(s, args...)
}

func (z *ZapLogger) With(args ...interface{}) {
	// TODO: Implement structured logging with fields using zap.Field
}

