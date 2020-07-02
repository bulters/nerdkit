package logging

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingReponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		request_id, ok := ctx.Value("requestID").(string)
		if !ok {
			request_id = "unknown"
		}
		logger := log.WithField("requestID", request_id)
		ctx = context.WithValue(ctx, "logger", logger)
		r = r.WithContext(ctx)

		time_start := time.Now()
		lrw := newLoggingReponseWriter(w)
		next.ServeHTTP(lrw, r)
		duration := time.Since(time_start)
		logger = logger.WithFields(log.Fields{
			"duration": duration,
			"status":   lrw.statusCode})
		logger.Info(r.RequestURI)
	})
}
