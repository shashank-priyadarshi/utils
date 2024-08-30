# utilities [![codecov](https://codecov.io/gh/shashank-priyadarshi/utils/graph/badge.svg?token=UL4Y1eBX07)](https://codecov.io/gh/shashank-priyadarshi/utils)

This repository contains common libraries for my projects.

## Unit test coverage graph

<img src="https://codecov.io/gh/shashank-priyadarshi/utils/graphs/tree.svg?token=UL4Y1eBX07" alt="Unit test coverage grid graph from Codecov"/>

This repository provides the following packages:

- [Algo](#algo)
- [Database](#database)
- [Logger](#logger)

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
