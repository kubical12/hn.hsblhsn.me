package caches

import "github.com/coocood/freecache"

var cache *freecache.Cache

func init() {
	cache = freecache.NewCache(50 * 1024 * 1024)
}

func Cache() *freecache.Cache {
	return cache
}
