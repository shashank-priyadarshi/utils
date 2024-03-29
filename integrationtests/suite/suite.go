package suite

import (
	"os"

	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/data"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/database"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/logger"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/network"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/pubsub"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/security"
	"github.com/shashank-priyadarshi/utilities/integrationtests/modules/worker"
)

const (
	DATA = iota
	DATABASE
	LOGGER
	NETWORK
	PUBSUB
	SECURITY
	WORKER
)

func Test() {

	var packages = make(map[int]string)
	var packageTests = make(map[int]func())

	// Read packages to be tested from environment variable
	packages[DATA] = os.Getenv("DATA")
	packages[DATABASE] = os.Getenv("DATABASE")
	packages[LOGGER] = os.Getenv("LOGGER")
	packages[NETWORK] = os.Getenv("NETWORK")
	packages[PUBSUB] = os.Getenv("PUBSUB")
	packages[SECURITY] = os.Getenv("SECURITY")
	packages[WORKER] = os.Getenv("WORKER")

	// Store package integration tests for execution
	packageTests[DATA] = data.Test
	packageTests[DATABASE] = database.Test
	packageTests[LOGGER] = logger.Test
	packageTests[NETWORK] = network.Test
	packageTests[PUBSUB] = pubsub.Test
	packageTests[SECURITY] = security.Test
	packageTests[WORKER] = worker.Test

	for key, value := range packages {
		switch value {
		case "true":
			packageTests[key]()
		}
	}
}
