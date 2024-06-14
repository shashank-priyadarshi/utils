package internal

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	// Customize Logrus configuration as needed (e.g., set formatter, output)
	return &LogrusLogger{
		logger: logger,
	}
}

func (l *LogrusLogger) Info(s string, args ...interface{}) {
	l.logger.Infof(s, args...)
}

func (l *LogrusLogger) Warn(s string, args ...interface{}) {
	l.logger.Warnf(s, args...)
}

func (l *LogrusLogger) Error(err error, args ...interface{}) {
	l.logger.Errorf("%v: %s", err, fmt.Sprint(args...))
}

func (l *LogrusLogger) Fatal(err error, args ...interface{}) {
	l.logger.Fatalf("%v: %s", err, fmt.Sprint(args...))
}

func (l *LogrusLogger) Debug(s string, args ...interface{}) {
	l.logger.Debugf(s, args...)
}

func (l *LogrusLogger) With(args ...interface{}) {
	// TODO: Implement structured logging with fields using logrus.Fields
}
