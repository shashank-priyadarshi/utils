# utilities

This repository contains common libraries for my projects.

This repository provides the following packages:

- [Logger](#logger)
- [Database](#database)
- [Network](#network)
- [Security](#security)

## Logger

```
type Logger interface {
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(error, ...interface{})
	Fatal(error, ...interface{})
	Debug(string, ...interface{})
	With(key string, args ...interface{})
}
```

## Database

```
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
```

## Network

## Security
