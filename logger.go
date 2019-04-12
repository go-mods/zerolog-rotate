package logger

import (
	"github.com/arthurkiller/rollingWriter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

// Aliases
type ZrRotateConfig = rollingwriter.Config
type ZrConsoleConfig = zerolog.ConsoleWriter
type ZrLogger = zerolog.Logger

// Config
type Config struct {
	// Log file options
	RwConfig func(*ZrRotateConfig)
	// Console writer options
	CwConfig func(*ZrConsoleConfig)
}

// New creates a root logger with file and console output
func New(options Config) ZrLogger {
	// Create new logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// Enable stack trace
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// Define file log
	writer := NewRotateWriter(options.RwConfig)
	// Set rotate writer as the global output
	logger = logger.Output(writer)
	// Add console writer hook
	hook := NewConsoleWriterHook(options.CwConfig)
	return logger.Hook(hook)
}
