package echo

import (
	"context"
	"fmt"
	"github.com/shashank-priyadarshi/utilities"
	"github.com/shashank-priyadarshi/utilities/network/models"
	"github.com/shashank-priyadarshi/utilities/network/ports/application"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server *echo.Echo
	config *models.Config
}

func New(config *models.Config) (*Server, error) {
	return &Server{config: config}, nil
}

func (e *Server) Start(args ...interface{}) (client application.Client, err error) {

	if len(args) < 1 {
		return nil, utilities.InsufficientParameters
	}

	var addr = fmt.Sprintf("%s:%d", e.config.Options.Host, e.config.Options.Port)

	server := echo.New()
	server.Use(middleware.Logger(), middleware.Recover(), middleware.RequestID())

	switch e.config.Options.TLS {
	case true:
		err = server.StartTLS(addr, e.config.Options.CertPath, e.config.Options.KeyPath)
	default:
		err = server.Start(addr)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %s", utilities.OperationFailed, err)
	}

	return &Client{
		server: e.server,
		Handler: Handler{
			server: e.server,
		},
	}, nil
}

func (e *Server) Shutdown(args ...interface{}) (err error) {

	var ctx context.Context

	if len(args) < 1 {
		ctx = context.TODO()
	}

	err = e.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", utilities.OperationFailed, err)
	}

	return nil
}

type Client struct {
	server *echo.Echo
	Handler
}

func (c *Client) Group(args ...interface{}) (application.Group, error) {

	group := &Group{
		server: c.server,
	}

	newEchoGroup, err := group.create(args...)
	if err != nil {
		return nil, err
	}

	group.group = newEchoGroup
	return group, nil
}

type Group struct {
	server *echo.Echo
	group  *echo.Group
}

func (g *Group) Group(args ...interface{}) (application.Group, error) {

	group, err := g.create(args...)
	if err != nil {
		return g, err
	}

	g.group = group

	return g, nil
}

func (g *Group) create(args ...interface{}) (*echo.Group, error) {

	argsLen := len(args)

	if argsLen < 1 {
		return nil, utilities.InsufficientParameters
	}

	var (
		prefix          string
		ok              bool
		inputFunc       func(...interface{}) error
		middlewareFunc  echo.MiddlewareFunc
		middlewareFuncs []echo.MiddlewareFunc
	)

	if prefix, ok = args[0].(string); !ok {
		return nil, utilities.InvalidParameter
	}

	if argsLen > 1 {
		for i := 1; i < argsLen; i++ {

			if inputFunc, ok = args[i].(func(...interface{}) error); !ok {
				return nil, utilities.InvalidParameter
			}

			middlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					return inputFunc(c)
				}
			}

			middlewareFuncs = append(middlewareFuncs, middlewareFunc)
		}
	}

	return g.server.Group(prefix, middlewareFuncs...), nil
}

func (g *Group) Handler(args ...interface{}) error {

	method, prefix, handlerFunc, middlewareFuncs, err := fetchHandlerArgs(args...)
	if err != nil {
		return err
	}

	g.group.Add(method, prefix, handlerFunc, middlewareFuncs...)

	return nil
}

type Handler struct {
	server *echo.Echo
}

func (h *Handler) Handler(args ...interface{}) error {

	method, prefix, handlerFunc, middlewareFuncs, err := fetchHandlerArgs(args...)
	if err != nil {
		return err
	}

	h.server.Add(method, prefix, handlerFunc, middlewareFuncs...)

	return nil
}

func fetchHandlerArgs(args ...interface{}) (string, string, echo.HandlerFunc, []echo.MiddlewareFunc, error) {
	argsLen := len(args)

	if argsLen < 3 {
		return "", "", nil, nil, utilities.InsufficientParameters
	}

	var (
		method, prefix  string
		ok              bool
		inputFunc       func(...interface{}) error
		handlerFunc     echo.HandlerFunc
		middlewareFunc  echo.MiddlewareFunc
		middlewareFuncs []echo.MiddlewareFunc
	)

	if method, ok = args[0].(string); !ok {
		return method, prefix, handlerFunc, middlewareFuncs, utilities.InvalidParameter
	}

	if prefix, ok = args[1].(string); !ok {
		return method, prefix, handlerFunc, middlewareFuncs, utilities.InvalidParameter
	}

	if inputFunc, ok = args[2].(func(...interface{}) error); !ok {
		return method, prefix, handlerFunc, middlewareFuncs, utilities.InvalidParameter
	}

	handlerFunc = func(c echo.Context) error {
		return inputFunc(c)
	}

	if argsLen > 3 {
		for i := 3; i < argsLen; i++ {

			if inputFunc, ok = args[i].(func(...interface{}) error); !ok {
				return method, prefix, handlerFunc, middlewareFuncs, utilities.InvalidParameter
			}

			middlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					return inputFunc(c)
				}
			}

			middlewareFuncs = append(middlewareFuncs, middlewareFunc)
		}
	}

	return method, prefix, handlerFunc, middlewareFuncs, nil
}
