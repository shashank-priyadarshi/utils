package logger

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/logger/constants"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

func NewLogger() (ports.Logger, error) {

	if !isSupported("logging") {
		return nil, fmt.Errorf("logging option %s is not supported", "")
	}

	return nil, nil
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
