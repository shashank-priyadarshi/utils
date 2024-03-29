package echo

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
	"github.com/shashank-priyadarshi/utilities/network/ports"
)

type EchoServer struct {
	log    loggerPort.Logger
	server echo.Echo
}

func NewEchoServer(log loggerPort.Logger) (*EchoServer, error) {
	return &EchoServer{log: log}, nil
}

func (e *EchoServer) Start(args ...interface{}) (client ports.Client, err error) {

	if len(args) < 1 {
		err = fmt.Errorf("")
		return
	}

	var (
		addr string
		ok   bool
	)

	if addr, ok = args[0].(string); !ok {
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

func (e *EchoServer) Shutdown(args ...interface{}) (err error) {
	var ctx context.Context

	if len(args) < 1 {
		ctx = context.TODO()
	}

	err = e.server.Shutdown(ctx)
	if err != nil {
		err = fmt.Errorf("%w", err)
	}

	return
}

type Group struct {
}

func (g *Group) Chain(i ...interface{}) ports.Group {
	return nil
}

func (g *Group) Add(i ...interface{}) error {
	return nil
}
