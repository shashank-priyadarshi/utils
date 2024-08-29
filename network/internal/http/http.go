package http

import (
	"go.ssnk.in/utils/network/ports"
	"go.ssnk.in/utils/network/types"
)

type Http struct {
	options *types.Options
}

func New(opts ...func(*Http)) (ports.Network, error) {
	http := &Http{}

	for _, opt := range opts {
		opt(http)
	}

	return http, nil
}

func WithConfig(options *types.Options) func(*Http) {
	return func(http *Http) {
		http.options = options
	}
}
