package caches

import "time"

func NewDurationMap() DurationMap {
	var (
		day  = time.Hour * 24
		week = day * 7
	)
	return DurationMap{
		".json": time.Hour,
		".html": day,
		".jpeg": week,
		".js":   week,
		".css":  week,
	}
}

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
