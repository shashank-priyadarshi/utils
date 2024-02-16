package rest

import (
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest/echo"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func NewRESTServer() (ports.REST, error) {
	return echo.NewEchoServer()
}
