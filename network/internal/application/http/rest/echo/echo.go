package echo

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
)

type EchoServer struct {
	log    loggerPort.Logger
	server echo.Echo
}

func NewEchoServer(log loggerPort.Logger) (*EchoServer, error) {
	return &EchoServer{log: log}, nil
}

func (e *EchoServer) AddGroup(i ...interface{}) error {
	return nil
}

func (e *EchoServer) AddHandler(i ...interface{}) error {
	return nil
}

func (e *EchoServer) AddMiddleware(i ...interface{}) error {
	return nil
}

func (e *EchoServer) Start(params ...interface{}) (err error) {
	if len(params) < 1 {
		err = fmt.Errorf("")
		return
	}

	var (
		addr string
		ok   bool
	)

	if addr, ok = params[0].(string); !ok {
		err = fmt.Errorf("")
		return
	}

	server := echo.New()
	server.Use(middleware.Logger(), middleware.Recover(), middleware.RequestID())

	err = server.Start(addr)
	if err != nil {
		err = fmt.Errorf("")
		return
	}

	return
}

func (e *EchoServer) Shutdown(params ...interface{}) (err error) {
	var ctx context.Context

	if len(params) < 1 {
		ctx = context.TODO()
	}

	err = e.server.Shutdown(ctx)
	if err != nil {
		err = fmt.Errorf("%w", err)
	}

	return
}
