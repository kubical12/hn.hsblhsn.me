package main

import (
	"github.com/hsblhsn/hn.hsblhsn.me/api"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityserver"
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

	logger.Info("main: starting server...")
	go readabilityserver.Initialize()
	if err := api.ListenAndServe(); err != nil {
		logger.Fatal("main: could not start server", zap.Error(err))
	}
}
