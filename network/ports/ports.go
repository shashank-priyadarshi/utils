package ports

import "io"

type Network interface {
	Dialector
	Server
	Closer
}

type Dialector interface {
	Dial() (io.ReadCloser, error)
}

type Server interface {
	Configure() (interface{}, error)
	Start() error
}

type Closer interface {
	Close() error
}
