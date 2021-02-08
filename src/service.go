package src

import (
	"github.com/cache/src/cacheImpl"
	"github.com/cache/src/manager"
)

type Service struct {
	Manager manager.Service
}

func NewService(maxSize int, policy string /* eviction instance*/) Service {

	mangerService := manager.NewService(cacheImpl.CacheDAO{}, maxSize, policy)

	return Service{
		Manager: *mangerService,
	}
}
