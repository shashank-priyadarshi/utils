package http

import (
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

type Server struct {
	ports.REST
	ports.GraphQL
}

func New() (ports.HTTP, error) {
	restServer, _ := rest.New()
	graphQLServer, _ := graphql.New()

	return &Server{
		restServer,
		graphQLServer,
	}, nil
}
