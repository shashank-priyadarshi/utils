package models

// Caller needs to specify the protocol and package he/she wants to use
type (
	Type string

	Protocol Type

	// DataLink layer protocol
	DataLink Protocol
	PPP      DataLink

	// Transport layer protocols
	Transport Protocol
	TCP       Transport
	UDP       Transport
	SCTP      Transport

	// Sessions layers protocols
	Sessions Protocol
	RPC      Sessions

	// Application layers protocols
	Application Protocol
	HTTP        Application
	FTP         Application
	SMTP        Application
	DNS         Application
	WS          Application
	REST        HTTP
	GRAPHQL     HTTP

	Config struct {
		Protocol struct {
			Layer                  Protocol
			Type, Name, Toolchaink Type
		}

		Options Options
	}

	Options struct {
	}
)
