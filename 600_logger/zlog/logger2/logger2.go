package zl

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Zlog *ZerologLogger

func Init() {
	Zlog = NewZerologLogger()
}

type ZerologLogger struct {
	logger *zerolog.Logger
}

func NewZerologLogger() *ZerologLogger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	output := zerolog.MultiLevelWriter(file, zerolog.ConsoleWriter{Out: os.Stdout})
	logger := zerolog.New(output).With().Timestamp().Caller().Logger()

	return &ZerologLogger{logger: &logger}
}

func (l *ZerologLogger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *ZerologLogger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *ZerologLogger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *ZerologLogger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *ZerologLogger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *ZerologLogger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

type Logger interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Panic() *zerolog.Event
}
