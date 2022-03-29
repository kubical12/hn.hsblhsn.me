package caches

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

func NewInMemoryCache() Cache {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic(err)
	}
	return cache
}
