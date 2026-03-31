package logger

import (
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var (
	Log  *zap.SugaredLogger
	once sync.Once
)

func Initialize() *zap.SugaredLogger {
	once.Do(func() {
		baseLogger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		Log = baseLogger.Sugar()
	})
	return Log
}

type (
	ResponseData struct {
		Status int
		Size   int
	}

	LoggingResponseWriter struct {
		http.ResponseWriter
		ResponseData *ResponseData
	}
)

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.ResponseData.Size += size
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.ResponseData.Status = statusCode
}
