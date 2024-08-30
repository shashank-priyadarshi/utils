package logger

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/logger"
	"os"
	"strconv"
)

func Test() {

	logProvider := os.Getenv("LOG_PROVIDER")
	logLevel := os.Getenv("LOG_LEVEL")
	logFormat := os.Getenv("LOG_FORMAT")
	traceRaw := os.Getenv("LOG_TRACE")

	trace, _ := strconv.ParseBool(traceRaw)

	log, err := logger.NewLogger(logProvider, logLevel, logFormat, trace)
	if err != nil {
		fmt.Println("failed to initialize new logger using provider ", logProvider)
	}

	log.With("key", "value") // debug
	log.Info("something", "info_key", "info_value")
	log.Debug("something", "debug_key", "debug_value") //debug
	log.Warn("something", "warn_key", "warn_value")
	log.Error(fmt.Errorf("some error"), "err_key", "err_value")
	log.Fatal(fmt.Errorf("some error"), "fatal_key", "fatal_value") //debug
}
