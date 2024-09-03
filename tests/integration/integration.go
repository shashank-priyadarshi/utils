package integration

import (
	"fmt"
	_ "github.com/ory/dockertest/v3"
	loggerPorts "go.ssnk.in/utils/logger/ports"
	"go.ssnk.in/utils/tests/integration/modules/algo"
	"go.ssnk.in/utils/tests/integration/modules/database"
	"go.ssnk.in/utils/tests/integration/modules/logger"
	"go.ssnk.in/utils/tests/types"
)

type Integration struct {
	logger   loggerPorts.Logger
	packages []types.Package
}

func New(logger loggerPorts.Logger) *Integration {
	return &Integration{
		logger: logger,
	}
}

func (i *Integration) Execute(c []types.Config) error {
	i.logger.Info("Setting up integration tests...")

	packageTests := make(map[types.Package]func())

	// Store package integration tests for execution
	packageTests[types.Algo] = algo.Test
	packageTests[types.Database] = database.Test
	packageTests[types.Logger] = logger.Test

	for _, pkg := range c {
		if _, ok := pkg.Config[types.Integration]; ok {
			i.logger.Info(fmt.Sprintf("Executing integration tests for package %d", pkg.Package))
			packageTests[pkg.Package]()
		}
	}

	return nil
}
