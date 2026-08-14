package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		duration := slog.Duration("time", time.Since(start))
		slog.Info("", "statusCode", wrapper.StatusCode, "method", r.Method, "path", r.URL.Path, "duration", duration)

	})
}
