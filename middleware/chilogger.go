package middleware

import (
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-mods/zerolog-rotate"
	"github.com/go-mods/zerolog-rotate/log"
)

// DefaultLogFormatter is a simple logger that implements a LogFormatter.
type DefaultLogger struct {
	Logger *logger.ZrLogger
}

// ChiLogger is a middleware for chi that logs the start and end of each request,
// along with some useful data about what was requested, what the response status
// was, and how long it took to return.
//
// It will use the main logger declared in zerolog-rotate.Logger
//
// example:
//	r := chi.NewRouter()
//	r.Use(middleware.ChiLogger)
func ChiLogger(next http.Handler) http.Handler {
	fn := middleware.RequestLogger(&DefaultLogger{&log.Logger})
	return fn(next)
}

// RequestLogger returns a logger handler using a custom zerolog-rotate.Logger.
//
// example:
//	r := chi.NewRouter()
//	r.Use(middleware.ChiLogger(logger))
func RequestLogger(logger *logger.ZrLogger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&DefaultLogger{logger})
}

// NewLogEntry creates a new LogEntry for the request.
func (l *DefaultLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &defaultLogEntry{
		InfoEvent:    l.Logger.Info(),
		DebugEvent:   l.Logger.Debug(),
		ErrorEvent:   l.Logger.Error(),
		errorMessage: "",
		error:        false,
	}

	infoMsg := ""

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		entry.DebugEvent.Str("request_id", reqID)
		entry.InfoEvent.Str("request_id", reqID)
		entry.ErrorEvent.Str("request_id", reqID)
		infoMsg = fmt.Sprintf("[%s] ", reqID)
		entry.errorMessage = fmt.Sprintf("[%s] ", reqID)
	}

	entry.InfoEvent.Str("request", "started")
	entry.DebugEvent.Str("request", "complete")
	entry.ErrorEvent.Str("request", "error")

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	entry.DebugEvent.Str("http_scheme", scheme)
	entry.DebugEvent.Str("http_proto", r.Proto)
	entry.DebugEvent.Str("http_method", r.Method)
	entry.DebugEvent.Str("remote_addr", r.RemoteAddr)
	entry.DebugEvent.Str("user_agent", r.UserAgent())

	uri := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
	entry.DebugEvent.Str("uri", uri)

	entry.InfoEvent.Msgf(infoMsg+"\\%s %s from %s", r.Method, uri, r.RemoteAddr)

	return entry
}

type defaultLogEntry struct {
	InfoEvent    *zerolog.Event
	DebugEvent   *zerolog.Event
	ErrorEvent   *zerolog.Event
	errorMessage string
	error        bool
}

// GetLogEntry returns the in-context LogEntry for a request.
func GetLogEntry(r *http.Request) *defaultLogEntry {
	entry := middleware.GetLogEntry(r).(*defaultLogEntry)
	return entry
}

// Add field to log entry
func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*defaultLogEntry); ok {
		entry.DebugEvent = entry.DebugEvent.Interface(key, value)
	}
}

// Add field to log entry
func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*defaultLogEntry); ok {
		entry.DebugEvent = entry.InfoEvent.Fields(fields)
	}
}

func (entry *defaultLogEntry) Write(status, bytes int, elapsed time.Duration) {
	entry.DebugEvent = entry.DebugEvent.Fields(map[string]interface{}{
		"response_status": status, "response_bytes_length": bytes,
		"response_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})
	entry.DebugEvent.Msg(entry.errorMessage)

	if entry.error {
		entry.ErrorEvent = entry.ErrorEvent.Fields(map[string]interface{}{
			"response_status": status, "response_bytes_length": bytes,
			"response_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
		})
		entry.ErrorEvent.Msg(entry.errorMessage)
	}
}

func (entry *defaultLogEntry) Panic(v interface{}, stack []byte) {
	entry.ErrorEvent = entry.ErrorEvent.Fields(map[string]interface{}{
		"stack": string(stack),
	})

	entry.error = true
	entry.errorMessage = entry.errorMessage + fmt.Sprintf("%+v", v)
}
