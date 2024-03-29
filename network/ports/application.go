package ports

type Application interface {
	HTTP
}

type HTTP interface {
	REST
	GraphQL
}

type GraphQL interface {
}
