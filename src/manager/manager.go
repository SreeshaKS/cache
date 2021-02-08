package manager

import (
	"container/list"
	"errors"

	"github.com/cache-implementation/src/cache"
)

type ValueData struct {
	Key  string
	Data interface{}
}

type Service struct {
	CacheDAO cache.ICacheA
	Map      map[string]*list.Element
	Queue    *list.List
}

func NewService(cacheDAO cache.ICacheA, maxSize int, policy string) *Service {
	//initialise eviction and add instance here
	return &Service{
		CacheDAO: cacheDAO,
		Map:      make(map[string]*list.Element, 3),
		Queue:    list.New(),
	}
}

func (s Service) Update(key string, cacheItem interface{}) error {
	return s.CacheDAO.Update(key, cacheItem)
}

func (s Service) Get(key string) (interface{}, error) {
	if _, ok := s.Map[key]; !ok {
		return nil, errors.New("key does not exist")
	}
	element := s.Map[key]
	s.Queue.MoveToFront(element)
	return element.Value, nil
}

func (s Service) Set(key string, item interface{}) error {
	valueData := ValueData{
		Key:  key,
		Data: item,
	}

	if _, ok := s.Map[key]; ok {
		s.Map[key].Value = valueData
		s.Queue.MoveToFront(s.Map[key])
		return nil
	}

	if len(s.Map) == 3 {
		element := s.Queue.Back()
		valueData, _ := element.Value.(ValueData)
		delete(s.Map, valueData.Key)
	}
	e := s.Queue.PushFront(valueData)
	s.Map[key] = e
	return nil
}

func (s Service) Items() (interface{}, error) {
	return s.CacheDAO.Items()
}
