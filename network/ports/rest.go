package ports

type REST interface {
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

//type Middleware interface {
//	Set(...interface{}) error
//}
