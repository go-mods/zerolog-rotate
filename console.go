package logger

import (
	"github.com/rs/zerolog"
)

type consoleWriterHook struct {
	writer zerolog.ConsoleWriter
	logger zerolog.Logger
}

// NewConsoleWriterHook creates and initializes a new consoleWriterHook.
func NewConsoleWriterHook(options ...func(w *zerolog.ConsoleWriter)) zerolog.Hook {
	// Create a new console writer
	writer := zerolog.NewConsoleWriter(options...)
	// Create a new logger
	logger := zerolog.New(writer).With().Timestamp().Logger()
	// Hook
	hook := consoleWriterHook{
		writer: writer,
		logger: logger,
	}
	return hook
}

// Implement zerolog.Hook interface
func (hook consoleWriterHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	switch level {
	case zerolog.InfoLevel:
		hook.logger.Info().Msg(message)
	case zerolog.WarnLevel:
		hook.logger.Warn().Msg(message)
	case zerolog.ErrorLevel:
		hook.logger.Error().Msg(message)
	case zerolog.NoLevel:
		hook.logger.Log().Msg(message)
	}
}
