package graphql

import (
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql/gqlgen"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New() (ports.GraphQL, error) {
	return gqlgen.New()
}
