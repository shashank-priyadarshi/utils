package ports

type Logger interface {
	Panic(error, ...interface{})
	Fatal(error, ...interface{})
	Error(error, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

type Sugar interface {
	With(map[string]string) Logger
}

type Leveler interface {
	Level(string) Logger
}
