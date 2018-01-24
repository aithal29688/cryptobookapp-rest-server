package server

import (
	"net/http"
)

type LoggedWriter struct {
	http.ResponseWriter

	statusCode int
}

func NewLoggedWriter(w http.ResponseWriter) *LoggedWriter {
	return &LoggedWriter{w, http.StatusOK}
}

func (w *LoggedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode

	w.ResponseWriter.WriteHeader(statusCode)
}

func (s *Server) WrapRequest(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Only handle requests that are not OPTIONS pre-flight ones
		lw := NewLoggedWriter(w)
		if r.Method != "OPTIONS" {
			inner.ServeHTTP(lw, r)
		}

	})
}
