package main

import (
	_ "github.com/ory/dockertest/v3"

	"github.com/shashank-priyadarshi/utilities/integrationtests/suite"
)

func main() {
	suite.Test()
}
