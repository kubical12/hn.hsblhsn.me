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
	io.Copy(resp.ResponseWriter, resp.output)
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
	cacheControlVal := fmt.Sprintf("public; max-age=%d; s-maxage=%d", cacheMaxAge, cacheMaxAge)
	resp.Header().Set("Cache-Control", cacheControlVal)
}

type DurationMap map[string]time.Duration

func (m DurationMap) Get(ext string, def time.Duration) int {
	dur, exists := m[ext]
	if !exists {
		dur = def
	}
	maxCacheAge := int(dur.Seconds())
	return maxCacheAge
}

type CacheOptions struct {
	Cache           Cache
	DefaultDuration time.Duration
	DurationMap     DurationMap
}

func NewDurationMap() DurationMap {
	return DurationMap{
		".json": time.Hour,
		".jpeg": time.Hour * 72,
		".html": time.Hour * 72,
		".js":   time.Hour * 24,
	}
}

func Middleware(hn http.Handler, opts CacheOptions) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		// do not do anything if it's not a get request.
		if r.Method != http.MethodGet {
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
			w.statusCode = http.StatusOK
			w.Write(cachedResp)
			return
		}

		// request not found in cache.
		// So, process the request.
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
