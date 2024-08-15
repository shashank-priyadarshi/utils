package logger

import (
	"fmt"

	"github.com/shashank-priyadarshi/utilities/logger/constants"
	"github.com/shashank-priyadarshi/utilities/logger/internal"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Logger struct {
	ports.Logger
}

func New(logProvider, logLevel, format string, trace bool) (Logger, error) {

	if !isSupported(logProvider) {
		return Logger{}, fmt.Errorf("logging option %s is not supported", "")
	}

	var log ports.Logger

	switch constants.Type(logProvider) {
	case constants.SLOG:
		log = internal.NewSlogLogger(logLevel, format, trace)
	case constants.LOGRUS:
	case constants.ZAP:
	case constants.ZEROLOG:
	}

	return Logger{log}, nil
}

func isSupported(option string) bool {
	var supported = make(map[constants.Type]any)

	supported[constants.SLOG] = nil
	supported[constants.LOGRUS] = nil
	supported[constants.ZAP] = nil
	supported[constants.ZEROLOG] = nil

	if _, ok := supported[constants.Type(option)]; ok {
		return true
	}

	return false
}
