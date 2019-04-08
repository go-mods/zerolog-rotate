package log

import (
	"context"
	"github.com/rs/zerolog"
)

// Logger is the global logger.
var Logger zerolog.Logger

// Debug starts a new message with debug level
// or sends the *Event with msg added as the message field if not empty.
func Debug(i ...interface{}) *zerolog.Event {
	if len(i) > 0 {
		for _, param := range i {
			if msg, ok := param.(string); ok {
				Logger.Debug().Msg(msg)
			}
		}
		return nil
	} else {
		return Logger.Debug()
	}
}

// Info starts a new message with info level
// or sends the *Event with msg added as the message field if not empty.
func Info(i ...interface{}) *zerolog.Event {
	if len(i) > 0 {
		for _, param := range i {
			if msg, ok := param.(string); ok {
				Logger.Info().Msg(msg)
			}
		}
		return nil
	}
	return Logger.Info()
}

// Warn starts a new message with warn level
// or sends the *Event with msg added as the message field if not empty.
func Warn(i ...interface{}) *zerolog.Event {
	if len(i) > 0 {
		for _, param := range i {
			if msg, ok := param.(string); ok {
				Logger.Warn().Msg(msg)
			}
		}
		return nil
	}
	return Logger.Warn()
}

// Error starts a new message with error level
// or sends the *Event with msg added as the message field if not empty.
func Error(i ...interface{}) *zerolog.Event {
	if len(i) == 1 {
		if err, ok := i[0].(error); ok {
			Logger.Error().Stack().Err(err).Msg(err.Error())
		} else if msg, ok := i[0].(string); ok {
			Logger.Error().Msg(msg)
		}
		return nil
	} else if len(i) == 2 {
		_, isError := i[0].(error)
		_, isString := i[1].(string)
		if isError && isString {
			Logger.Error().Stack().Err(i[0].(error)).Msg(i[1].(string))
		}
	}
	return Logger.Error()
}

// Fatal starts a new message with fatal level
// or sends the *Event with msg added as the message field if not empty.
func Fatal(i ...interface{}) *zerolog.Event {
	if len(i) > 0 {
		for _, param := range i {
			if msg, ok := param.(string); ok {
				Logger.Fatal().Msg(msg)
			}
		}
		return nil
	}
	return Logger.Fatal()
}

// Panic starts a new message with panic level
// or sends the *Event with msg added as the message field if not empty.
func Panic(i ...interface{}) *zerolog.Event {
	if len(i) > 0 {
		for _, param := range i {
			if msg, ok := param.(string); ok {
				Logger.Panic().Msg(msg)
			}
		}
		return nil
	}
	return Logger.Panic()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
