package rest

import (
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http/rest/echo"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

func NewRESTServer(log loggerPort.Logger) (ports.REST, error) {
	return echo.NewEchoServer(log)
}
