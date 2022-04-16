package main

import (
	"context"
	"log"
	"net/http"

	"github.com/blendle/zapdriver"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend"
	"github.com/hsblhsn/hn.hsblhsn.me/frontend"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	var (
		client = &http.Client{Transport: http.DefaultTransport}
		router = mux.NewRouter()
	)
	app := fx.New(
		fx.Supply(client),
		fx.Supply(router),
		fx.Provide(newLogger),
		backend.Module(),
		frontend.Module(),
		fx.Invoke(startHTTPServer),
	)
	app.Run()
}

func newLogger() (*zap.Logger, error) {
	logger, err := zapdriver.NewProduction()
	if err != nil {
		return nil, errors.Wrap(err, "main: failed to build logger")
	}
	return logger, nil
}

func startHTTPServer(lc fx.Lifecycle, router *mux.Router, logger *zap.Logger) {
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				logger.Info("starting http server", zap.String("addr", ":8080"))
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatal("main: error starting server", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(ctx); err != nil {
				return errors.Wrap(err, "main:error shutting down server")
			}
			return nil
		},
	})
}
