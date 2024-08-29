package secretxform

import (
	"go.ssnk.in/utils/secretxform/constants"
	"go.ssnk.in/utils/secretxform/models"
	"go.ssnk.in/utils/secretxform/ports"
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
