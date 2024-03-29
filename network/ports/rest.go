package ports

type REST interface {
	Start(...interface{}) (Client, error)
	Shutdown(...interface{}) error
}

type Client interface {
	Group
	Handler
}

type Group interface {
	Chain(...interface{}) Group
	Handler
}

type Handler interface {
	Add(...interface{}) error
}

//type Middleware interface {
//	Set(...interface{}) error
//}
