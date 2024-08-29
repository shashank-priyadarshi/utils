package types

import "net/http"

type (
	// Layer: Datalink, Transport, Sessions, Application
	// Standard: REST, GraphQL
	// Library:  Standard // gRPC, Echo, gqlgen

	// Protocol : PPP, TCP, UDP, SCTP, RPC, HTTP, FTP, SMTP, DNS, WS
	Protocol string

	Network struct {
		Protocol Protocol
		Options  *Options
	}

	Options struct {
		CertPath, KeyPath string
		Host              string
		Port              uint16
	}
)

type (
	HTTPMiddlewareFunc func(handler http.Handler) http.Handler
)
