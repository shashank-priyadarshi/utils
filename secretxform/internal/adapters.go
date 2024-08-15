package adapters

import (
	"github.com/shashank-priyadarshi/utilities/secretxform/constants"
	"github.com/shashank-priyadarshi/utilities/secretxform/internal/jwt"
	"github.com/shashank-priyadarshi/utilities/secretxform/models"
	"github.com/shashank-priyadarshi/utilities/secretxform/ports"
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
