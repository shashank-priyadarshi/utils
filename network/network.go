package network

import (
	"go.ssnk.in/utils/network/constants"
	"go.ssnk.in/utils/network/internal/http"
	"go.ssnk.in/utils/network/ports"
	"go.ssnk.in/utils/network/types"
)

func New(opts ...func(*types.Network)) (ports.Network, error) {

	n := &types.Network{}

	for _, opt := range opts {
		opt(n)
	}

	switch n.Protocol {
	case constants.HTTP:
		return http.New(http.WithConfig(n.Options))
	default:
		return nil, utilities.UnsupportedType
	}
}

func WithProtocol(protocol string) func(*types.Network) {
	return func(network *types.Network) {
		network.Protocol = types.Protocol(protocol)
	}
}

func WithConfig(options *types.Options) func(*types.Network) {
	return func(network *types.Network) {
		network.Options = options
	}
}
