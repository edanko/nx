package logger

import (
	"fmt"
	"net/http"
	"time"

	httpmiddleware "github.com/edanko/nx/pkg/http/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

// StructuredLogger is wrapper for zerolog
type StructuredLogger struct {
	Logger zerolog.Logger
}

// NewStructuredLogger use to init logger middleware
func NewStructuredLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

// NewLogEntry implement logging when receive request
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{
		Logger: l.Logger.Info(),
	}

	if rec := recover(); rec != nil {
		entry = &StructuredLoggerEntry{
			Logger: l.Logger.Error(),
		}
	}

	logFields := map[string]any{}

	logFields["req_time"] = time.Now()

	if reqID := httpmiddleware.GetRequestID(r.Context()); reqID != "" {
		logFields["request_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	if referer := r.Referer(); referer != "" {
		logFields["referer"] = referer
	}

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.Fields(logFields)

	return entry
}

// StructuredLoggerEntry is wrapper for zerolog.Event
type StructuredLoggerEntry struct {
	Logger *zerolog.Event
}

// Write is method that was call when server response the request
func (l *StructuredLoggerEntry) Write(
	status, bytes int,
	_ http.Header,
	elapsed time.Duration,
	_ any,
) {
	l.Logger = l.Logger.Fields(map[string]any{
		"resp_status":       status,
		"resp_bytes_length": bytes,
		"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1_000_000.0,
	})
	l.Logger.Msg("request completed")
}

// Panic is method that was call when server have panic with request
func (l *StructuredLoggerEntry) Panic(v any, stack []byte) {
	l.Logger = l.Logger.Fields(map[string]any{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
	l.Logger.Msg("request failed")
}
