package servers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
)

func Serve(h http.Handler) error {
	srv := server.New(h, &server.Options{
		RequestLogger: requestlog.NewStackdriverLogger(os.Stdout, func(error) {}),
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := srv.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		return errors.Wrap(err, "api: failed to start server")
	}
	return nil
}
