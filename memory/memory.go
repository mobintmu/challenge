package memory

import "hash/fnv"

type data struct {
	key       string
	value     string
	timestamp int64
}

type InMemoryStorage struct {
	buckets  [][]*data
	size     int
	lastUsed []*data
}

func NewInMemoryStorage(maximumCapacity int, maxSize int) *InMemoryStorage {
	return &InMemoryStorage{
		buckets:  make([][]*data, maximumCapacity),
		size:     0,
		lastUsed: make([]*data, 0, maxSize),
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
		println(index, " : ", value.key)
	}
	println("<<<<<<<<<<<<<<<<<<<")

}

func (s *InMemoryStorage) 
