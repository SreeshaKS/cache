package eviction

import "github.com/cache-implementation/src/cache"

type interface IEviction {
	func Evict()
}

type Service {
	// Policy
	Policy string
	// Cache Instance
	CacheDAO cache.ICacheA
	Eviction IEviction 
}

func NewService(string policy, cacheInstance cache.ICacheA, eviction IEviction) {
	return Service {
		Policy: policy,
		CacheDAO: cacheInstance,
		Eviction: eviction,
	}
}
func (s Service) Evict() {
	s.Eviction.Evict()
}