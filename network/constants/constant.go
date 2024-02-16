package constants

import "github.com/shashank-priyadarshi/utilities/network/models"

const (
	DataLink    models.Protocol = "datalink"
	Transport   models.Protocol = "transport"
	Sessions    models.Protocol = "sessions"
	Application models.Protocol = "application"

	PPP models.DataLink = "ppp"

	TCP  models.Transport = "tcp"
	UDP  models.Transport = "udp"
	SCTP models.Transport = "sctp"

	HTTP models.Application = "http"
	FTP  models.Application = "ftp"
	SMTP models.Application = "smtp"
	DNS  models.Application = "dns"
	WS   models.Application = "ws"

	GRPC   models.RPC     = "grpc"
	ECHO   models.REST    = "echo"
	GQLGEN models.GRAPHQL = "gqlgen"
)
