package ports

type DataLink interface {
}

type Server interface {
	Start(...interface{}) (Client, error)
	Shutdown(...interface{}) error
}
