package main

import (
	"fmt"
	"github.com/go-mods/zerolog-rotate"
	"github.com/go-mods/zerolog-rotate/log"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"time"
)

func main() {

	// Create the main logger
	log.Logger = logger.New(logger.Config{
		RwConfig: func(rw *logger.ZrRotateConfig) {
			rw.LogPath = "examples/logs"
			rw.FileName = "main"
			rw.TimeTagFormat = time.RFC3339
		},
		CwConfig: func(cw *logger.ZrConsoleConfig) {
			cw.NoColor = true
			cw.TimeFormat = time.StampMilli
			cw.FormatLevel = func(i interface{}) string {
				return fmt.Sprintf("| %-6s|", strcase.ToCamel(i.(string)))
			}
		},
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
