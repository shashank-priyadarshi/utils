package ports

type Protocol interface {
	DataLink
	Transport
	Sessions
	Application
}

// Connection interface might be removed altogether
type Connection interface {
	Connect(...interface{}) error
	Disconnect(...interface{}) error
}
