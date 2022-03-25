package main

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/servers"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/spa"
	"github.com/hsblhsn/hn.hsblhsn.me/embedded"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()
	zap.ReplaceGlobals(logger)
	logger.Info("main: starting ui server...")
	router := mux.NewRouter()
	spa.RegisterRoutes(router, embedded.Assets)
	if err := servers.Serve(router); err != nil {
		logger.Fatal("main: could not start server", zap.Error(err))
	}
}
