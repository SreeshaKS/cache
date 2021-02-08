package cache

type Item interface{}

type ICacheA interface {
	Update(key string, cacheItem Item) error
	Get(key string) (Item, error)
	Set(key string, item Item) error
	Items() ([]Item, error)
}
