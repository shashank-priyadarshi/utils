package models

type (
	Type     string
	Layer    Type     // Datalink, Transport, Sessions, Application
	Protocol Layer    // PPP, TCP, UDP, SCTP, RPC, HTTP, FTP, SMTP, DNS, WS
	Standard Protocol // REST, GraphQL
	Library  Standard // gRPC, Echo, gqlgen

	Config struct {
		Network Network
		Options Options
	}

	Network struct {
		Layer    Layer
		Protocol Protocol
		Standard Standard
		Library  Library
	}
	Options struct {
	}
)
