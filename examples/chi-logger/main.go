package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-mods/zerolog-rotate"
	"github.com/go-mods/zerolog-rotate/log"
	zrmiddleware "github.com/go-mods/zerolog-rotate/middleware"
	"github.com/iancoleman/strcase"
	"net/http"
	"time"
)

func main() {

	// Create the main logger
	log.Logger = logger.New(logger.Config{
		RwConfig: func(rw *logger.ZrRotateConfig) {
			rw.LogPath = "examples/logs"
			rw.FileName = "chi"
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

	// Routes
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(zrmiddleware.ChiLogger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})
	r.Get("/wait", func(w http.ResponseWriter, r *http.Request) {
		zrmiddleware.LogEntrySetField(r, "wait", true)
		time.Sleep(1 * time.Second)
		_, _ = w.Write([]byte("hi"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oops")
	})

	_ = http.ListenAndServe(":3333", r)
}
