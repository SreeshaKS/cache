package eviction

import "github.com/cache/src/cache"

type IEviction interface {
	Evict()
}

type Service struct {
	// Policy
	Policy string
	// Cache Instance
	CacheDAO cache.ICacheA
	Eviction IEviction
}

func NewService(policy string, cacheInstance cache.ICacheA, eviction IEviction) *Service {
	return &Service{
		Policy:   policy,
		CacheDAO: cacheInstance,
		Eviction: eviction,
	}
}
func (s Service) Evict() {
	s.Eviction.Evict()
}
