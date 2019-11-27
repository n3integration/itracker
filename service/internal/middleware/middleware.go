package middleware

import (
	"github.com/n3integration/itracker/internal/logger"
	"net/http"
)

type statusWriter struct {
	code     int
	delegate http.ResponseWriter
}

func (s statusWriter) Code() int {
	if s.code == 0 {
		return http.StatusOK
	}
	return s.code
}

func (s statusWriter) Header() http.Header {
	return s.delegate.Header()
}

func (s statusWriter) Write(b []byte) (int, error) {
	return s.delegate.Write(b)
}

func (s *statusWriter) WriteHeader(statusCode int) {
	s.code = statusCode
	s.delegate.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		writer := &statusWriter{delegate: w}
		defer logger.Info("%s %s %d", req.Method, req.URL.Path, writer.Code())
		next.ServeHTTP(writer, req)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST")
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		next.ServeHTTP(w, req)
	})
}
