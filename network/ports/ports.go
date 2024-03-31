package ports

// Network is the main interface that all network types should implement
type Network interface{}

// Protocol is the interface that all protocols should implement
type Protocol interface{}

// Standard is the interface that all standards should implement
type Standard interface{}

// Library is the interface that all libraries should implement
type Library interface{}

// Connection interface might be removed altogether
//type Connection interface {
//	Connect(...interface{}) error
//	Disconnect(...interface{}) error
//}
