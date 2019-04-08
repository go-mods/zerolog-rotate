package main

import (
	"github.com/pkg/errors"
	"github.com/go-mods/zerolog-rotate"
	"github.com/go-mods/zerolog-rotate/log"
)

func main() {

	// Create the main logger
	log.Logger = logger.New("example/logs", "example")

	// Info
	log.Info().Str("foo", "bar").Msg("You can use zerolog-rotate like you would do with zerolog")
	log.Info().Msg("Log your info like zerolog")
	log.Info("This is a quickest way to log")
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
