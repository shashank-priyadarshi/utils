package integration

import (
	"fmt"
	_ "github.com/ory/dockertest/v3"
	"go.ssnk.in/utils/test/integration/modules/algo"
	"go.ssnk.in/utils/test/integration/modules/database"
	"go.ssnk.in/utils/test/integration/modules/logger"
	"go.ssnk.in/utils/test/types"
)

type Integration struct {
	packages []types.Package
}

func New() *Integration {
	return &Integration{}
}

func (i *Integration) Execute(c *types.Config) error {
	fmt.Println("Setting up integration tests...")

	var packageTests = make(map[types.Package]func())

	// Store package integration tests for execution
	packageTests[types.Algo] = algo.Test
	packageTests[types.Database] = database.Test
	packageTests[types.Logger] = logger.Test

	fmt.Println("Executing integration tests...")

	for _, value := range i.packages {
		packageTests[value]()
	}

	return nil
}
