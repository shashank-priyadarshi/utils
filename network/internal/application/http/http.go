package http

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/graphql"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New(config *models.Config) (ports.Standard, error) {

	switch config.Network.Standard {
	case constants.REST:
		return rest.New(config)
	case constants.GRAPHQL:
		return graphql.New(config)
	default:
		return nil, utilities.UnsupportedType

	}
}
