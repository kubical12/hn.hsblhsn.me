package caches

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

const (
	DefaultCacheExpiration = 10 * time.Minute
)

var _ Cache = (*bigcache.BigCache)(nil)

func NewInMemoryCache() *bigcache.BigCache {
	cfg := bigcache.DefaultConfig(DefaultCacheExpiration)
	cache, err := bigcache.NewBigCache(cfg)
	if err != nil {
		panic(err)
	}
	return cache
}
