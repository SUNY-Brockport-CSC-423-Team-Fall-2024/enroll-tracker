package middleware

import (
	"enroll-tracker/internal/models"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

// LoggingMiddleware logs the incoming HTTP request, it's duration, and errors if there is a panic()
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		//Middleware logic here
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error("error occured",
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()
			//Start timer to time it takes to complete next.ServeHTTP
			start := time.Now()
			//Wrapped response writer to capture response status
			wrapped := models.WrapResponseWriter(w)
			//Call the next handler
			next.ServeHTTP(wrapped, r)
			//Log after request has been served
			logger.Info("request received",
				"status", wrapped.Status(),
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
				"src_ip", r.RemoteAddr,
			)
		}

		return http.HandlerFunc(fn)
	}
}
