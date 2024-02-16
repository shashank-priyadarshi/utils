package application

import (
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
	"github.com/shashank-priyadarshi/utilities/network/internal/application/http"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

type ApplicationProtocolServer struct {
	log loggerPort.Logger
	ports.Application
}

func NewApplicationProtocolServer(log loggerPort.Logger) (ports.Application, error) {
	appProto, _ := http.NewHTTPServer(log)

	return &ApplicationProtocolServer{log, appProto}, nil
}
