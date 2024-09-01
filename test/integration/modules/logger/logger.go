package logger

import (
	"fmt"
	"os"
	"strconv"

	"go.ssnk.in/utils/logger"
)

func Test() {

	logProvider := os.Getenv("LOG_PROVIDER")
	logLevel := os.Getenv("LOG_LEVEL")
	logFormat := os.Getenv("LOG_FORMAT")
	traceRaw := os.Getenv("LOG_TRACE")

	trace, _ := strconv.ParseBool(traceRaw)

	log := logger.New(logProvider, logLevel, logFormat, trace)

	log.With("key", "value") // debug
	log.Info("something", "info_key", "info_value")
	log.Debug("something", "debug_key", "debug_value") //debug
	log.Warn("something", "warn_key", "warn_value")
	log.Error(fmt.Errorf("some error"), "err_key", "err_value")
	log.Fatal(fmt.Errorf("some error"), "fatal_key", "fatal_value") //debug
}
