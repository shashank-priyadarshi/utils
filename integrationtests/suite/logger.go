package suite

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/logger"
	"strconv"
)

func (t *testSuite) testLogger(args []string) {
	logProvider, logLevel, logFormat := args[0], args[1], args[2]
	trace, _ := strconv.ParseBool(args[3])

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
