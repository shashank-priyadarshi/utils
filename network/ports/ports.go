package ports

import "github.com/shashank-priyadarshi/utilities/network/ports/application"

// Network is the main interface that all network types should implement
type Network interface {
	Protocol
}

// Protocol is the interface that all protocols should implement
type Protocol interface {
	application.HTTP
}

// Connection interface might be removed altogether
//type Connection interface {
//	Connect(...interface{}) error
//	Disconnect(...interface{}) error
//}
