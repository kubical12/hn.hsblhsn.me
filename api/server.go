package api

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/hackernews"
)

// RegisterRoutes starts the server.
func RegisterRoutes(router *mux.Router) {
	imgProxyEndpoint := "/api/v1/feed_images"
	handler := NewHandler(hackernews.NewHackerNews(imgProxyEndpoint))
	handler.RegisterRoutes(router.PathPrefix("/api/v1").Subrouter())
}
