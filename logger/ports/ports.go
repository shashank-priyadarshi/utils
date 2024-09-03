package ports

type Logger interface {
	Dispatcher
	Sugar
	Leveler
}

type Dispatcher interface {
	Panic(error, ...interface{})
	Fatal(error, ...interface{})
	Error(error, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

type Sugar interface {
	With(map[string]string) Dispatcher
}

type Leveler interface {
	Level(string) Dispatcher
}
