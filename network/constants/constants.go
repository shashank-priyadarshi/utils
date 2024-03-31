package constants

import "github.com/shashank-priyadarshi/utilities/network/models"

const (
	DataLink    models.Layer = "datalink"
	Transport   models.Layer = "transport"
	Sessions    models.Layer = "sessions"
	Application models.Layer = "application"

	PPP models.Protocol = "ppp" // DataLink layer

	TCP  models.Protocol = "tcp" // Transport layer
	UDP  models.Protocol = "udp"
	SCTP models.Protocol = "sctp"

	RPC models.Protocol = "rpc" // Sessions layer

	HTTP models.Protocol = "http" // Application layer
	FTP  models.Protocol = "ftp"
	SMTP models.Protocol = "smtp"
	DNS  models.Protocol = "dns"
	WS   models.Protocol = "ws"

	REST    models.Standard = "rest"
	GRAPHQL models.Standard = "graphql"

	GRPC   models.Library = "grpc"
	ECHO   models.Library = "echo"
	GQLGEN models.Library = "gqlgen"
)
