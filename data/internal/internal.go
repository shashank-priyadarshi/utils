package internal

import (
	"github.com/shashank-priyadarshi/utilities/data/constants"
	"github.com/shashank-priyadarshi/utilities/data/internal/jwt"
	"github.com/shashank-priyadarshi/utilities/data/models"
	"github.com/shashank-priyadarshi/utilities/data/ports"
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
)

func NewDataHandler(logger loggerPort.Logger, config *models.Config) (ports.Data, error) {

	switch config.Data {
	case constants.JWT:
		return jwt.NewJWTHandler(logger)

	case constants.SAML:
		return nil, nil

	}

	return nil, nil
}
