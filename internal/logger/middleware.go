package logger

import (
	"net/http"
	"time"
)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseData := &ResponseData{}
		lw := LoggingResponseWriter{
			ResponseWriter: w,
			ResponseData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)
		if responseData.Status == 0 {
			responseData.Status = http.StatusOK
		}

		Initialize().Infow(
			"http request processed",
			"uri", r.RequestURI,
			"method", r.Method,
			"status", responseData.Status,
			"duration", duration,
			"size", responseData.Size,
		)
	}

	return http.HandlerFunc(logFn)
}
