package internal

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logrus struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *Logrus {
	logger := logrus.New()
	// Customize Logrus configuration as needed (e.g., set formatter, output)
	return &Logrus{
		logger: logger,
	}
}

func (l *Logrus) Info(s string, args ...interface{}) {
	l.logger.Infof(s, args...)
}

func (l *Logrus) Warn(s string, args ...interface{}) {
	l.logger.Warnf(s, args...)
}

func (l *Logrus) Error(err error, args ...interface{}) {
	l.logger.Errorf("%v: %s", err, fmt.Sprint(args...))
}

func (l *Logrus) Fatal(err error, args ...interface{}) {
	l.logger.Fatalf("%v: %s", err, fmt.Sprint(args...))
}

func (l *Logrus) Debug(s string, args ...interface{}) {
	l.logger.Debugf(s, args...)
}

func (l *Logrus) With(args ...interface{}) {
	// TODO: Implement structured logging with fields using logrus.Fields
}
