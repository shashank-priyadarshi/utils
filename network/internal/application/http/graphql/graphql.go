package graphql

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql/gqlgen"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New(config *models.Config) (ports.GraphQL, error) {

	switch config.Network.Library {
	case constants.GQLGEN:
		return gqlgen.New(config)
	default:
		return nil, utilities.UnsupportedType
	}
}
