package rest

import (
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest/echo"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func New() (ports.REST, error) {
	return echo.New()
}
