package integration

import (
	"fmt"
	_ "github.com/ory/dockertest/v3"
	"go.ssnk.in/utils/test/integration/modules/database"
	"go.ssnk.in/utils/test/integration/modules/logger"
	"go.ssnk.in/utils/test/integration/modules/network"
	"go.ssnk.in/utils/test/integration/modules/pubsub"
	"go.ssnk.in/utils/test/integration/modules/secretxform"
	"go.ssnk.in/utils/test/integration/modules/security"
	"go.ssnk.in/utils/test/integration/modules/worker"
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
	packageTests[types.DATA] = secretxform.Test
	packageTests[types.DATABASE] = database.Test
	packageTests[types.LOGGER] = logger.Test
	packageTests[types.NETWORK] = network.Test
	packageTests[types.PUBSUB] = pubsub.Test
	packageTests[types.SECURITY] = security.Test
	packageTests[types.WORKER] = worker.Test

	fmt.Println("Executing integration tests...")
	packageTests[types.WORKER]()

	for _, value := range i.packages {
		packageTests[value]()
	}

	return nil
}
