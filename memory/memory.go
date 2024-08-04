package memory

import (
	"hash/fnv"
	"time"
)

type data struct {
	key       string
	value     string
	timestamp int64
}

type InMemoryStorage struct {
	buckets    []*data
	size       int
	lastUsed   []*data
	maxSizeLRU int
}

func NewInMemoryStorage(maximumCapacity int, maxSizeLRU int) *InMemoryStorage {
	return &InMemoryStorage{
		buckets:    make([]*data, 0, maximumCapacity),
		size:       0,
		lastUsed:   make([]*data, 0, maxSizeLRU),
		maxSizeLRU: maxSizeLRU,
	}
}

func (s *InMemoryStorage) hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (s *InMemoryStorage) updateLRU(currentData *data) {

	for index, value := range s.lastUsed {
		if value.key == currentData.key {
			copy(s.lastUsed[index:], s.lastUsed[index+1:])
			s.lastUsed = s.lastUsed[:len(s.lastUsed)-1]
			break
		}
	}
	s.lastUsed = append([]*data{currentData}, s.lastUsed...)
}

func (s *InMemoryStorage) printLastUsed() {
	println(">>>>>>>>>>>>>>>>>>>")
	for index, value := range s.lastUsed {
		if value != nil {
			println(index, " : ", value.key)
		}
	}
	println("<<<<<<<<<<<<<<<<<<<")

}

func (s *InMemoryStorage) printBuckets() {
	println(">>>>buckets>>>>>>>>>>>>>")
	for index, value := range s.buckets {
		if value != nil {
			println(index, " : ", value.key)
		}
	}
	println("<<<<buckets<<<<<<<<<<<<")
}

// remove last item ( least recently used ) from the cache
func (s *InMemoryStorage) evict() {
	if len(s.lastUsed) == 0 {
		return
	}

	lru := s.lastUsed[len(s.lastUsed)-1]

	for index, value := range s.buckets {
		if value != nil {
			if value.key == lru.key {
				s.buckets = append(s.buckets[:index], s.buckets[index+1:]...)
				break
			}
		}
	}

	s.lastUsed = s.lastUsed[:len(s.lastUsed)-1]
	s.size--
}

func (s *InMemoryStorage) Set(key, value string) {

	//if exist update
	for index, value_bucket := range s.buckets {
		if value_bucket.key == key {
			s.buckets[index].value = value
			s.buckets[index].timestamp = time.Now().UnixNano()
			s.updateLRU(s.buckets[index])
			return
		}
	}

	if s.size >= s.maxSizeLRU {
		s.evict()
	}

	// add new entity
	newData := &data{
		key:       key,
		value:     value,
		timestamp: time.Now().UnixNano(),
	}
	s.buckets = append(s.buckets, newData)
	s.size++
	s.updateLRU(newData)
}
