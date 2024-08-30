package internal

import (
	"log/slog"
	"os"
)

const (
	LevelFatal = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelFatal: "FTL",
}

type Slog struct {
	logger *slog.Logger
}

func NewSlogLogger(logLevel, format string, trace bool) *Slog {
	logger := &Slog{}
	logger.init(logLevel, format, trace)
	return logger
}

func (s *Slog) init(logLevel, format string, trace bool) {
	var opts = &slog.HandlerOptions{}
	var handler slog.Handler

	opts.AddSource = trace
	opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		//if a.Key == slog.TimeKey {
		//}

		if a.Key == slog.LevelKey {
			level := a.Value.Any().(slog.Level)
			levelLabel, exists := LevelNames[level]
			if exists {
				levelLabel = level.String()
			} else {
				switch level {
				case slog.LevelError:
					levelLabel = "ERR"
				case slog.LevelDebug:
					levelLabel = "DBG"
				case slog.LevelWarn:
					levelLabel = "WRN"
				default:
					levelLabel = "INF"
				}
			}
			a.Value = slog.StringValue(levelLabel)
		}

		return a
	}

	switch logLevel {
	case "error":
		opts.Level = slog.LevelError
	case "fatal":
		opts.Level = LevelFatal
	case "debug":
		opts.Level = slog.LevelDebug
	case "warn":
		opts.AddSource = false
		opts.Level = slog.LevelWarn
	default:
		opts.AddSource = false
		opts.Level = slog.LevelInfo
	}

	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	s.logger = slog.New(handler)
}

func (s *Slog) Info(msg string, args ...interface{}) {
	s.logger.Info(msg, args...)
}

func (s *Slog) Warn(msg string, args ...interface{}) {
	s.logger.Warn(msg, args...)
}

func (s *Slog) Error(err error, args ...interface{}) {
	s.logger.Error(err.Error(), args...)
}

func (s *Slog) Fatal(err error, args ...interface{}) {

}

func (s *Slog) Debug(msg string, args ...interface{}) {
	s.logger.Debug(msg, args...)
}

func (s *Slog) With(args ...interface{}) {
	var attrs []slog.Attr

	if len(args)/2 != 0 {
		return
	}

	var (
		ok    bool
		key   string
		value string
	)

	for i := 0; i < len(args)-1; i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return
		}

		value, ok = args[i+1].(string) // process value type
		if !ok {
			return
		}

		attr := slog.String(key, value)
		attrs = append(attrs, attr)
	}

	s.logger = s.logger.With(slog.Group(key, attrs))
}
