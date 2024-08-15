package secretxform

import (
	"github.com/shashank-priyadarshi/utilities/secretxform/constants"
	"github.com/shashank-priyadarshi/utilities/secretxform/internal"
	"github.com/shashank-priyadarshi/utilities/secretxform/models"
	"github.com/shashank-priyadarshi/utilities/secretxform/ports"
)

func New(config *models.Config) (ports.Data, error) {

	if !isSupported(config.Type) {
		return nil, nil
	}

	return adapters.New(config)
}

func isSupported(data models.Type) bool {
	switch data {
	case constants.JWT, constants.SAML:
		return true
	}
	return false
}
