package load

import (
	loggerPorts "go.ssnk.in/utils/logger/ports"
	"go.ssnk.in/utils/tests/types"
)

type Load struct {
	logger loggerPorts.Logger
}

func New(logger loggerPorts.Logger) *Load {
	return &Load{
		logger: logger,
	}
}

func (l *Load) Execute(c []types.Config) error {
	return nil
}
