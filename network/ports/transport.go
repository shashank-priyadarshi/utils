package ports

type Transport interface {
	TCP
	UDP
	SCTP
}

type TCP interface {
}

type UDP interface {
}

type SCTP interface {
}
