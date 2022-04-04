package caches

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"time"
)

type ResponseWriter struct {
	statusCode int
	output     *bytes.Buffer
	http.ResponseWriter
}

var _ http.ResponseWriter = (*ResponseWriter)(nil)

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		statusCode:     0,
		output:         bytes.NewBuffer(nil),
		ResponseWriter: w,
	}
}

func (resp *ResponseWriter) WriteHeader(code int) {
	resp.statusCode = code
}

func (resp *ResponseWriter) Write(p []byte) (n int, err error) {
	return resp.output.Write(p)
}

func (resp *ResponseWriter) flush() {
	resp.ResponseWriter.WriteHeader(resp.statusCode)
	_, _ = io.Copy(resp.ResponseWriter, resp.output)
}

func (resp *ResponseWriter) process(r *http.Request, m DurationMap, def time.Duration) {
	defer resp.flush()
	ext := path.Ext(r.URL.Path)
	resp.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'")
	resp.Header().Set("X-XSS-Protection", "1; mode=block")
	resp.Header().Set("X-Frame-Options", "DENY")
	resp.Header().Set("Content-Type", mime.TypeByExtension(ext))

	if resp.statusCode != http.StatusOK {
		return
	}
	cacheMaxAge := m.Get(ext, def)
	cacheControlVal := fmt.Sprintf("public, max-age=%d", cacheMaxAge)
	resp.Header().Set("Cache-Control", cacheControlVal)
}

func Middleware(hn http.Handler, opts CacheOptions) http.Handler {
	const cacheStatusHeader = "X-Cache-Status"
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		// do not do anything if it's not a get request.
		if r.Method != http.MethodGet {
			resp.Header().Set(cacheStatusHeader, "DYNAMIC")
			hn.ServeHTTP(resp, r)
			return
		}

		// request is cacheable.
		var (
			cacheKey = r.RequestURI
			w        = NewResponseWriter(resp)
		)
		defer w.process(r, opts.DurationMap, opts.DefaultDuration)

		// look for response in the cache.
		cachedResp, err := opts.Cache.Get(cacheKey)
		if err == nil {
			// found in cache. send cached response.
			w.Header().Set(cacheStatusHeader, "HIT")
			w.statusCode = http.StatusOK
			_, _ = w.Write(cachedResp)
			return
		}

		// request not found in cache.
		// So, process the request.
		w.Header().Set(cacheStatusHeader, "MISS")
		hn.ServeHTTP(w, r)
		if w.statusCode == 0 {
			w.statusCode = http.StatusOK
		}
		// do not cache any unsuccessful response.
		if w.statusCode != http.StatusOK {
			return
		}
		if err := opts.Cache.Set(cacheKey, w.output.Bytes()); err != nil {
			log.Println(err)
		}
	})
}
