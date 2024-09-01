package internal

import (
	"fmt"

	"go.uber.org/zap"
)

type Zap struct {
	logger *zap.Logger
}

func NewZapLogger() *Zap {
	logger, _ := zap.NewProduction() // Replace with your desired Zap configuration
	return &Zap{
		logger: logger,
	}
}

func (z *Zap) Info(s string, args ...interface{}) {
	z.logger.Sugar().Infof(s, args...)
}

func (z *Zap) Warn(s string, args ...interface{}) {
	z.logger.Sugar().Warnf(s, args...)
}

func (z *Zap) Error(err error, args ...interface{}) {
	z.logger.Sugar().Errorf("%v: %s", err, fmt.Sprint(args...))
}

func (z *Zap) Fatal(err error, args ...interface{}) {
	z.logger.Sugar().Fatalf("%v: %s", err, fmt.Sprint(args...))
}

func (z *Zap) Debug(s string, args ...interface{}) {
	z.logger.Sugar().Debugf(s, args...)
}

func (z *Zap) With(args ...interface{}) {
	// TODO: Implement structured logging with fields using zap.Field
}
