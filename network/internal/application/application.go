package application

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New(config *models.Config) (ports.Protocol, error) {

	switch config.Network.Protocol {
	case constants.HTTP:
		return http.New(config)
	case constants.FTP:
		return nil, nil
	case constants.SMTP:
		return nil, nil
	case constants.DNS:
		return nil, nil
	case constants.WS:
		return nil, nil
	default:
		return nil, utilities.UnsupportedType
	}
}
