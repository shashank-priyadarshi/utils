package echo

import loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"

type EchoServer struct {
	log loggerPort.Logger
}

func NewEchoServer() (*EchoServer, error) {
	return nil, nil
}

func (e *EchoServer) AddGroup(i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (e *EchoServer) AddHandler(i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (e *EchoServer) AddMiddleware(i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (e *EchoServer) Start(i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (e *EchoServer) Shutdown(i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}
