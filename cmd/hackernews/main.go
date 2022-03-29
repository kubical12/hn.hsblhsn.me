package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend"
	"github.com/hsblhsn/hn.hsblhsn.me/caches"
	"github.com/hsblhsn/hn.hsblhsn.me/frontend"
)

func main() {
	var (
		router = mux.NewRouter()
		apiV1  = router.PathPrefix("/api/v1").Subrouter()
		root   = router.PathPrefix("/").Subrouter()
	)

	frontend.RegisterRoutes(root)
	backend.RegisterRoutes(apiV1)

	handler := caches.Middleware(router, caches.CacheOptions{
		Cache:           caches.NewInMemoryCache(),
		DefaultDuration: time.Hour,
		DurationMap:     caches.NewDurationMap(),
	})
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatal(err)
	}
}
