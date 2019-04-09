package main

import (
	"github.com/go-mods/zerolog-rotate"
	"github.com/go-mods/zerolog-rotate/log"
	"github.com/pkg/errors"
	"time"
)

func main() {

	// Create the main logger
	log.Logger = logger.New(logger.Config{
		RwConfig: func(rw *logger.RotateConfig) { rw.LogPath = "example/logs"; rw.FileName = "example" },
		CwConfig: func(cw *logger.ConsoleConfig) { cw.TimeFormat = time.RFC3339 },
	})

	// Info
	log.Info().Str("foo", "bar").Msg("You can use zerolog-rotate like you would do with zerolog")
	log.Info().Msg("Log your message like zerolog")
	log.Info("Or use this way to quickly log")
	log.Debug("Debug logs only appear in log file")

	// Error
	err := errors.Wrap(errors.New("error message"), "from error")
	log.Error().Stack().Err(err).Msg("You can log errors in different ways")
	log.Error(err)
	log.Error(err, "Append a message to your error")
	log.Error("This is a custom error")

	// Print
	log.Print("Print is like Debug")
}
