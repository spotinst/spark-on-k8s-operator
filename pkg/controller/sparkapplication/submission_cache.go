package sparkapplication

/*
A simple cache implementation that stores keys with a time-to-live (TTL) value.
*/

import (
	"github.com/jellydator/ttlcache/v3"
	"time"
)

type SubmissionCache struct {
	cache *ttlcache.Cache[string, any] // value is not used
}

func NewSubmissionCache(ttl time.Duration) *SubmissionCache {
	cache := ttlcache.New[string, any](
		ttlcache.WithTTL[string, any](ttl),
	)

	c := &SubmissionCache{
		cache: cache,
	}
	go cache.Start() // start the cache cleanup goroutine
	return c
}

func (c *SubmissionCache) Set(key string) {
	c.cache.Set(key, nil, ttlcache.DefaultTTL)
}

func (c *SubmissionCache) Exist(key string) bool {
	return c.cache.Has(key)
}
