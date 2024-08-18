package integration

import (
	"fmt"
	_ "github.com/ory/dockertest/v3"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/database"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/logger"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/network"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/pubsub"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/secretxform"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/security"
	"github.com/shashank-priyadarshi/utilities/test/integration/modules/worker"
	"github.com/shashank-priyadarshi/utilities/test/types"
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
