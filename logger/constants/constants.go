package constants

type Type string

const (
	SLOG    Type = "slog"
	ZAP     Type = "zap"
	LOGRUS  Type = "logrus"
	ZEROLOG Type = "zerolog"
)
