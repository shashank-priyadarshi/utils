package constants

type Provider string

const (
	Slog    Provider = "slog"
	Zap     Provider = "zap"
	Logrus  Provider = "logrus"
	Zerolog Provider = "zerolog"
)

func (p Provider) String() string {
	return string(p)
}

type Level string

const (
	Panic Level = "panic"
	Fatal Level = "fatal"
	Error Level = "error"
	Warn  Level = "warn"
	Info  Level = "info"
	Debug Level = "debug"
)

func (l Level) String() string {
	return string(l)
}

type Format string

const (
	Json  Format = "json"
	Text  Format = "text"
	Proto Format = "proto"
)

func (f Format) String() string {
	return string(f)
}
