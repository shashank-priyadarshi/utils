package profile

import (
	loggerPorts "go.ssnk.in/utils/logger/ports"
	"go.ssnk.in/utils/tests/types"
)

type Profile struct {
	logger loggerPorts.Logger
}

func New(logger loggerPorts.Logger) *Profile {
	return &Profile{
		logger: logger,
	}
}

func (m *Profile) Execute(c []types.Config) error {
	return nil
}
