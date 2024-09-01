package models

import (
	"go.ssnk.in/utils/database/constants"
)

type (
	Config struct {
		Options Options
		Type    constants.Database
	}

	Options struct {
		LogOptions
		DatabaseOptions
	}

	LogOptions struct {
		Provider, Level, Format string
		Trace                   bool
	}

	DatabaseOptions struct {
		URI, Username, Password string
		Driver                  constants.Driver
		ORM                     constants.ORM
		WithORM                 bool
		MongoDBPoolSize         uint8
	}

	Response struct {
		Result []interface{}
	}
)
