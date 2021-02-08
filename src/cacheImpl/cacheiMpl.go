package cacheImpl

import (
	"sync"

	"github.com/cache-implementation/src/cache"
)

type CacheDAO struct {
	Lock *sync.RWMutex
}

func (s CacheDAO) Update(key string, cacheItem cache.Item) error {
	return nil
}

func (s CacheDAO) Get(key string) (cache.Item, error) {
	return nil, nil
}

func (s CacheDAO) Set(key string, item cache.Item) error {
	return nil
}

func (s CacheDAO) Items() ([]cache.Item, error) {
	return nil, nil
}
