package logger

import (
	"fmt"
	"os"

	"github.com/Oleska1601/WBOptimizeServer/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

type LoggerI interface {
	Debug() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
}

var _ LoggerI = (*Logger)(nil)

func New(cfg *config.LoggerConfig) (*Logger, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logLevel, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("parse level: %w", err)
	}

	logger = logger.Level(logLevel)
	return &Logger{
		logger: &logger,
	}, nil
}

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
