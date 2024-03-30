package adapters

import (
	"github.com/shashank-priyadarshi/utilities/data/constants"
	"github.com/shashank-priyadarshi/utilities/data/internal/jwt"
	"github.com/shashank-priyadarshi/utilities/data/models"
	"github.com/shashank-priyadarshi/utilities/data/ports"
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
