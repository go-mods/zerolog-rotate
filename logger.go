package logger

import (
	"github.com/arthurkiller/rollingWriter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

// New creates a root logger with file and console output
func New(LogPath string, FileName string) zerolog.Logger {
	// Create new logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// Enable stack trace
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// Define file log
	writer := NewRotateWriter(
		func(rc *rollingwriter.Config) {
			rc.LogPath = LogPath
			rc.FileName = FileName
		},
	)
	// Set rotate writer as the global output
	logger = logger.Output(writer)
	// Add console writer hook
	hook := NewConsoleWriterHook(
		func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.RFC3339
		}, )
	return logger.Hook(hook)
}