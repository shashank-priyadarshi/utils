package network

type Type string

type Protocol Type

const (
	TCP   Protocol = "tcp"
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
	WS    Protocol = "ws"
	WSS   Protocol = "wss"
)

type RPC Type

const (
	GRPC RPC = "grpc"
	TRPC RPC = "trpc"
)

type REST Type

const (
	GIN      REST = "gin"
	MUX      REST = "mux"
	FASTHTTP REST = "fasthttp"
	ECHO     REST = "echo"
)
