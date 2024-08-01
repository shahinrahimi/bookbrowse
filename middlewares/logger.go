package middlewares

import "net/http"

func (m *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		m.logger.Printf("Request %s to %s from %s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(rw, r)
	})
}
