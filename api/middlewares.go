package api

import (
	"fmt"
	"net/http"
	"time"
)

// AttachMiddlewares attaches necessary middlewares to the given handler.
func AttachMiddlewares(fn http.HandlerFunc, reqTimeout, cacheTimeout time.Duration) http.Handler {
	timeoutMsg := ErrorMsg{Message: "Request timed out"}
	h := http.TimeoutHandler(fn, reqTimeout, timeoutMsg.String())
	return CORS(CacheControl(h, cacheTimeout))
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
