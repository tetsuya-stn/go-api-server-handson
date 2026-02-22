package middlewares

import (
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceId := newTraceId()
		ctx := SetTraceId(req.Context(), traceId)
		req = req.WithContext(ctx)
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, req)

		log.Printf("[%d]%s %s %d\n", traceId, req.Method, req.RequestURI, rw.statusCode)
	})
}
