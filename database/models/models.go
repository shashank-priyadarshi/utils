package models

import (
	"github.com/shashank-priyadarshi/utilities/database/constants"
)

type Config struct {
	Options Options
	Type    constants.Database
}

type Options struct {
	LogOptions
	DatabaseOptions
}

type LogOptions struct {
	Provider, Level, Format string
	Trace                   bool
}

type DatabaseOptions struct {
	URI, Username, Password string
	Driver                  constants.Driver
	ORM                     constants.ORM
	WithORM                 bool
	MongoDBPoolSize         uint8
}

type Response struct {
	Error  error
	Result []interface{}
}
