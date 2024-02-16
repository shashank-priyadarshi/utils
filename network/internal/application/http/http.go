package http

import (
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

type HTTPServer struct {
	log loggerPort.Logger
	ports.REST
	ports.GraphQL
}

func NewHTTPServer(log loggerPort.Logger) (ports.HTTP, error) {
	restServer, _ := rest.NewRESTServer()
	graphQLServer, _ := graphql.NewGraphQLServer()

	return &HTTPServer{
		log,
		restServer,
		graphQLServer,
	}, nil
}
