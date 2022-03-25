package main

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityserver"
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
	logger.Info("main: starting api server...")
	go readabilityserver.Initialize()

	// Start the server.
	router := mux.NewRouter()
	spa.RegisterRoutes(router.PathPrefix("/").Subrouter(), embedded.Assets)
	if err := servers.Start(router); err != nil {
		logger.Fatal("main: could not start server", zap.Error(err))
	}
}
