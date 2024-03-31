package rest

import (
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/constants"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest/echo"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports/application"
)

func New(config *models.Config) (application.HTTP, error) {

	switch config.Network.Library {
	case constants.ECHO:
		return echo.New()
	default:
		return nil, utilities.UnsupportedType
	}
}
