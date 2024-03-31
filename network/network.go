package network

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New(config *models.Config) (ports.Network, error) {

	switch config.Network.Layer {
	case constants.Application:
		return application.New(config)
	default:
		return nil, utilities.UnsupportedType
	}
}
