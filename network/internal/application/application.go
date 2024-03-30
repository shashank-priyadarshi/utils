package application

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

type Server struct {
	ports.Application
}

func New(config *models.Config) (ports.Application, error) {

	switch models.Application(config.Protocol.Type) {
	case constants.HTTP:
		return http.New()
	default:
		return nil, utilities.UnsupportedType
	}
}
