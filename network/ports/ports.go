package ports

type Protocol interface {
	Application
}

type Application interface {
	HTTP
}

type HTTP interface {
	REST
	GraphQL
}

type REST interface {
	AddGroup(...interface{}) error
	AddHandler(...interface{}) error
	AddMiddleware(...interface{}) error
	Server
}

type GraphQL interface {
}

type Server interface {
	Start(...interface{}) error
	Shutdown(...interface{}) error
}

// Connection interface might be removed altogether
type Connection interface {
	Connect(...interface{}) error
	Disconnect(...interface{}) error
}
