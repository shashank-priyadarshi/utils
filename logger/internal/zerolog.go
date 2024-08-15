package internal

import (
	"github.com/rs/zerolog"
)

type Zerolog struct {
	logger zerolog.Logger
}

func NewZerolog() *Zerolog {
	return &Zerolog{
		logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (z *Zerolog) Info(s string, args ...interface{}) {
	z.logger.Info().Msgf(s, args...)
}

func (z *Zerolog) Warn(s string, args ...interface{}) {
	z.logger.Warn().Msgf(s, args...)
}

func (z *Zerolog) Error(err error, args ...interface{}) {
	z.logger.Error().Err(err).Msgf(fmt.Sprint(args...))
}

func (z *Zerolog) Fatal(err error, args ...interface{}) {
	z.logger.Fatal().Err(err).Msgf(fmt.Sprint(args...))
}

func (z *Zerolog) Debug(s string, args ...interface{}) {
	z.logger.Debug().Msgf(s, args...)
}

func (z *Zerolog) With(args ...interface{}) {
	// TODO: Implement structured logging with fields
}
