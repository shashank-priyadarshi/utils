package application

type HTTP interface {
	Start(...interface{}) (Client, error)
	Shutdown(...interface{}) error
}

type Client interface {
	Group(...interface{}) (Group, error)
	Handler
}

type Group interface {
	Group(...interface{}) (Group, error)
	Handler
}

type Handler interface {
	Handler(...interface{}) error
}
