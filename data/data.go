package data

import (
	"github.com/shashank-priyadarshi/utilities/data/constants"
	"github.com/shashank-priyadarshi/utilities/data/internal"
	"github.com/shashank-priyadarshi/utilities/data/models"
	"github.com/shashank-priyadarshi/utilities/data/ports"
)

func New(config *models.Config) (ports.Data, error) {

	if !isSupported(config.Data) {
		return nil, nil
	}

	return adapters.New(config)
}

func isSupported(data models.Data) bool {
	switch data {
	case constants.JWT, constants.SAML:
		return true
	}
	return true
}
