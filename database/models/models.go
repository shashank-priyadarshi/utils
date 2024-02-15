package models

import (
	"github.com/shashank-priyadarshi/utilities/database/constants"
)

type Config struct {
	Options Options
	Type    constants.Database
}

type Options struct {
	URI, Username, Password string
	Driver                  constants.Driver
	ORM                     constants.ORM
	WithORM                 bool
}

type Response struct {
	Error  error
	Result interface{}
}
