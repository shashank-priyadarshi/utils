package logger

import (
	"fmt"

	"go.ssnk.in/utils/logger"
)

func Test() {
	log := logger.New(logger.SetProvider("slog"), logger.SetLevel("warn"), logger.SetFormat("json"), logger.WithTracing())

	defer log.Fatal(fmt.Errorf("fatal error"), "fatal_key", "fatal_value") // debug
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Errorf("panic from log.Panic recovered"), "panic_error", r)
		}
	}()

	log.Info("info message", "info_key", "info_value")
	log.Debug("debug message", "debug_key", "debug_value") // not working
	log.Warn("warn message", "warn_key", "warn_value")
	log.Error(fmt.Errorf("error"), "err_key", "err_value")
	log.Panic(fmt.Errorf("panic error"), "panic_key", "panic_value") // panic log can be better
}
