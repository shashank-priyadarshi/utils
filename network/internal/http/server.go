package http

import (
	"go.ssnk.in/utils/network/ports"
	"go.ssnk.in/utils/network/types"
	_ "golang.org/x/net/websocket"
	"net/http"
)

func (h *Http) Configure() (interface{}, error) {
	return newServer(), nil
}

func (h *Http) Start() error {
	return nil
}

func (h *Http) Close() error {
	return nil
}

type server struct {
	*http.ServeMux
	ports.Group
	ports.Handler
	ports.WebSocket
}

func newServer() *server {
	s := http.NewServeMux()

	h := &handler{s}
	g := &group{groups: make(map[string][]types.HTTPMiddlewareFunc), h: h}
	ws := &ws{s}

	return &server{
		s, g, h, ws,
	}
}

type group struct {
	// TODO: this will be a trie to search for existing groups
	groups map[string][]types.HTTPMiddlewareFunc
	h      ports.Handler
}

func (g *group) Group(pattern string, middlewares ...types.HTTPMiddlewareFunc) (ports.Group, error) {

	var (
		//ms []types.HTTPMiddlewareFunc
		ok bool
	)

	if _, ok = g.groups[pattern]; ok {
		return g, utilities.NewError("group already exists")
	}

	g.groups[pattern] = middlewares

	return g, nil
}

func (g *group) Handler(pattern string, handler http.Handler, middlewares ...types.HTTPMiddlewareFunc) error {
	return g.h.Handler(pattern, handler, middlewares...)
}

type handler struct {
	*http.ServeMux
}

func (h *handler) Handler(pattern string, handler http.Handler, middlewares ...types.HTTPMiddlewareFunc) error {

	for _, m := range middlewares {
		handler = m(handler)
	}

	h.HandleFunc(pattern, h.ServeHTTP)

	return nil
}

type ws struct {
	*http.ServeMux
}

func (s *ws) Upgrade() error {
	return nil
}
