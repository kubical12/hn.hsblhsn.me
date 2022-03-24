package main

import (
	"log"

	"github.com/hsblhsn/hn.hsblhsn.me/api"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityserver"
)

func main() {
	go readabilityserver.Initialize()
	if err := api.ListenAndServe(); err != nil {
		log.Fatal("main: could not start server:", err)
	}
}
