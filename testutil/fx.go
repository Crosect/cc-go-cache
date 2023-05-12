package ccCacheTestUtil

import (
	"github.com/crosect/cc-go-cache"
	"go.uber.org/fx"
)

var cache *cccache.Cache

func EnableCacheTestUtil() fx.Option {
	return fx.Invoke(func(c *cccache.Cache) {
		cache = c
	})
}

// Cache return cache instance
func Cache() *cccache.Cache {
	return cache
}
