package servers

import (
	"net/http"
	"time"
)

func globalMiddlewares(h http.Handler) http.Handler {
	msg := `{"error":true,"message":"Request time out"}`
	return http.TimeoutHandler(csp(xss(h)), time.Second*10, msg)
}

func csp(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'")
		h.ServeHTTP(w, r)
	})
}

func xss(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "DENY")
		h.ServeHTTP(w, r)
	})
}
