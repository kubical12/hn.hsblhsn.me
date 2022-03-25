package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/caches"
)

// AttachMiddlewares attaches necessary middlewares to the given handler.
func AttachMiddlewares(fn http.HandlerFunc, reqTimeout, cacheTimeout time.Duration) http.Handler {
	timeoutMsg := ErrorMsg{Message: "Request timed out"}
	h := http.TimeoutHandler(fn, reqTimeout, timeoutMsg.String())
	return CSP(CORS(caches.Middleware(h, cacheTimeout)))
}

// CORS adds CORS headers to the response.
func CORS(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		if r.Method == "OPTIONS" {
			return
		}
		fn.ServeHTTP(w, r)
	})
}

// CacheControl adds cache control headers to the response.
func CacheControl(h http.Handler, d time.Duration) http.Handler {
	secs := d.Seconds()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int(secs)))
		h.ServeHTTP(w, r)
	})
}

// CSP adds Content-Security-Policy headers to the response.
func CSP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csp := "default-src 'self'; report-uri /api/v1/_/csp-reports"
		w.Header().Set("Content-Security-Policy", csp)
		h.ServeHTTP(w, r)
	})
}
