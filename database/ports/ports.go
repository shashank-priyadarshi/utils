package ports

import (
	"context"
	"github.com/shashank-priyadarshi/utilities/database/models"
)

type Database interface {
	Operations
	Transactions
	Closer
	Configure
}

type Operations interface {
	Create(context.Context, ...interface{}) (models.Response, error)
	Query(context.Context, ...interface{}) (models.Response, error)
	Update(context.Context, ...interface{}) (models.Response, error)
	Delete(context.Context, ...interface{}) (models.Response, error)
}

type Transactions interface {
	Begin(context.Context, ...interface{}) (models.Response, error)
	Execute(context.Context, ...interface{}) (models.Response, error)
	Rollback(context.Context, ...interface{}) (models.Response, error)
}

type Closer interface {
	Close() error
}

type Configure interface {
	Configure(context.Context, ...interface{}) error
}
