package caches

import "time"

type DurationMap map[string]time.Duration

func (m DurationMap) Get(ext string, def time.Duration) int {
	dur, exists := m[ext]
	if !exists {
		dur = def
	}
	maxCacheAge := int(dur.Seconds())
	return maxCacheAge
}

type CacheOptions struct {
	Cache           Cache
	DefaultDuration time.Duration
	DurationMap     DurationMap
}

func NewDurationMap() DurationMap {
	return DurationMap{
		".json": time.Hour,
		".jpeg": time.Hour * 72,
		".html": time.Hour * 72,
		".js":   time.Hour * 24,
	}
}
