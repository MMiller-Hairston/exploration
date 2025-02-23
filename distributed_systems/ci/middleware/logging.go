package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type recorder struct {
	http.ResponseWriter
	statusCode int
}

func (rw *recorder) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := recorder{w, 200}
			next.ServeHTTP(&rec, r)
			logger.Info("Incoming request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("host", r.URL.Host),
				slog.Int("code", rec.statusCode),
				slog.String("duration", time.Since(start).String()),
				slog.String("ua", r.UserAgent()),
			)
		})
	}
}
