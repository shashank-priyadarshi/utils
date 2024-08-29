package adapters

import (
	"go.ssnk.in/utils/secretxform/constants"
	"go.ssnk.in/utils/secretxform/internal/jwt"
	"go.ssnk.in/utils/secretxform/models"
	"go.ssnk.in/utils/secretxform/ports"
)

func New(config *models.Config) (ports.Data, error) {

	switch config.Type {
	case constants.JWT:
		return jwt.Handle(config)

	case constants.SAML:
		return nil, nil

	}

	return nil, nil
}
