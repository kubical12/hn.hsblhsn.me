package api

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/hackernews"
	"github.com/pkg/errors"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
)

// ListenAndServe starts the server.
func ListenAndServe() error {
	router := mux.NewRouter()
	imgProxyEndpoint := "http://localhost:8080/api/v1/feed_images"
	handler := NewHandler(hackernews.NewHackerNews(imgProxyEndpoint))
	handler.RegisterRoutes(router.PathPrefix("/api/v1").Subrouter())

	srv := server.New(router, &server.Options{
		RequestLogger: requestlog.NewStackdriverLogger(os.Stdout, func(error) {}),
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := srv.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port)); err != nil {
		return errors.Wrap(err, "api: failed to start server")
	}
	return nil
}
