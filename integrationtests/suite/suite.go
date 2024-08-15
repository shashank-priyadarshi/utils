package suite

import (
	"fmt"
	"os"
	"sync"

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

	fmt.Println("Setting up integration tests...")

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
	//packageTests[DATA] = secretxform.Test
	//packageTests[DATABASE] = database.Test
	//packageTests[LOGGER] = logger.Test
	//packageTests[NETWORK] = network.Test
	//packageTests[PUBSUB] = pubsub.Test
	//packageTests[SECURITY] = security.Test
	packageTests[WORKER] = worker.Test

	fmt.Println("Executing integration tests...")
	var wg sync.WaitGroup
	wg.Add(1)
	packageTests[WORKER]()

	//for key, value := range packages {
	//	switch {
	//	case value == "true":
	//		packageTests[key]()
	//	}
	//}

	wg.Wait()
}
