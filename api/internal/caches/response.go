package caches

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type CachedResponseWriter struct {
	statusCode int
	response   *bytes.Buffer
	http.ResponseWriter
}

func NewCachedResponseWriter(w http.ResponseWriter) *CachedResponseWriter {
	return &CachedResponseWriter{
		ResponseWriter: w,
		response:       bytes.NewBuffer(nil),
	}
}

func (w *CachedResponseWriter) Write(b []byte) (int, error) {
	return w.response.Write(b)
}

func (w *CachedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *CachedResponseWriter) send(secs int) {
	w.WriteHeader(w.statusCode)
	if w.statusCode == http.StatusOK {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", secs))
	} else {
		w.Header().Set("Cache-Control", "no-cache")
	}
	_, _ = io.Copy(w.ResponseWriter, w.response)
}

func Middleware(fn http.Handler, cacheTimeout time.Duration) http.Handler {
	secs := int(cacheTimeout.Seconds())
	cache := Cache()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewCachedResponseWriter(w)
		defer rw.send(secs)

		cacheKey := []byte("server-handler" + r.RequestURI)
		if r.Method == http.MethodGet {
			if item, err := cache.Get(cacheKey); err == nil {
				w.Header().Set("X-Cache", "HIT")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(item)
				return
			}
		}

		w.Header().Set("X-Cache", "MISS")
		fn.ServeHTTP(rw, r)
		if r.Method == http.MethodGet && rw.statusCode == http.StatusOK {
			err := cache.Set(cacheKey, rw.response.Bytes(), secs)
			if err != nil {
				zap.L().Error("caches: failed to cache handler response", zap.Error(err))
			}
		}
	})
}
