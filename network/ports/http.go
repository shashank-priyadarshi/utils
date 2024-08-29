package ports

import (
	"go.ssnk.in/utils/network/types"
	"net/http"
)

type Http interface {
	Group
	Handler
	WebSocket
}

type WebSocket interface {
	Upgrade() error
}

type Group interface {
	Group(string, ...types.HTTPMiddlewareFunc) (Group, error)
	Handler
}

type Handler interface {
	Handler(string, http.Handler, ...types.HTTPMiddlewareFunc) error
}
