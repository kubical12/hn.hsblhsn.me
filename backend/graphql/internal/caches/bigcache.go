package caches

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

const (
	CacheExpiration = 10 * time.Minute
)

var _ Cache = (*bigcache.BigCache)(nil)

func NewInMemoryCache() *bigcache.BigCache {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(CacheExpiration))
	if err != nil {
		panic(err)
	}
	return cache
}
