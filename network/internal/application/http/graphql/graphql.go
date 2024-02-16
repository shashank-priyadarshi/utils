package graphql

import (
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql/gqlgen"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func NewGraphQLServer() (ports.GraphQL, error) {
	return gqlgen.NewGraphQLServer()
}
