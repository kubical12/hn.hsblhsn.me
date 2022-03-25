package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/caches"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/hackernews"
)

const (
	DefaultReqTimeout  = time.Second * 10
	DefaultCacheMaxAge = time.Minute * 30
)

// cacheResp attaches cache middleware to the given handler.
func cacheResp(fn http.HandlerFunc, cacheTimeout time.Duration) http.Handler {
	return caches.Middleware(http.HandlerFunc(fn), cacheTimeout)
}

func setRequestTimeout(r *http.Request, d time.Duration) (*http.Request, context.CancelFunc) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, d)
	return r.WithContext(ctx), cancel
}

// RegisterRoutes starts the server.
func RegisterRoutes(router *mux.Router) {
	imgProxyEndpoint := "/api/v1/feed_images"
	handler := NewHandler(hackernews.NewHackerNews(imgProxyEndpoint))
	handler.RegisterRoutes(router)
}
