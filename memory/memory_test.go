package memory

import (
	"testing"
	"time"
)

func TestCreateHash(t *testing.T) {
	memory := NewInMemoryStorage(5, 100)
	hash := memory.hash("hi")
	if hash != 1748694682 {
		t.Errorf("Hash function is not working as expected")
	}
}

func TestAddUpdateLRU(t *testing.T) {
	A := &data{
		key:       "Ahi",
		value:     "1212",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	B := &data{
		key:       "Bbye",
		value:     "1312",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	C := &data{
		key:       "Chello",
		value:     "1234",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	myMemory := NewInMemoryStorage(5, 100)
	myMemory.updateLRU(A)
	myMemory.updateLRU(B)
	myMemory.updateLRU(C)
	myMemory.updateLRU(A)
	myMemory.updateLRU(A)
	myMemory.updateLRU(C)
	myMemory.updateLRU(A)

	// myMemory.printLastUsed()

	if myMemory.lastUsed[0].key != "Ahi" && myMemory.lastUsed[1].key != "Chello" && myMemory.lastUsed[2].key != "Bbye" {
		t.Errorf("LRU is not working as expected")
	}

}

func TestRemoveEvict(t *testing.T) {

	A := &data{
		key:       "Ahi",
		value:     "1212",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	B := &data{
		key:       "Bbye",
		value:     "1312",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	C := &data{
		key:       "Chello",
		value:     "1234",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	myMemory := NewInMemoryStorage(5, 100)
	myMemory.updateLRU(A)
	myMemory.updateLRU(B)
	myMemory.updateLRU(C)
	myMemory.updateLRU(A)
	myMemory.updateLRU(A)
	myMemory.updateLRU(C)
	myMemory.updateLRU(A)

	myMemory.buckets = append(myMemory.buckets, A)
	myMemory.buckets = append(myMemory.buckets, B)
	myMemory.buckets = append(myMemory.buckets, C)

	// myMemory.printBuckets()

	myMemory.evict()

	// myMemory.printBuckets()

	if myMemory.buckets[0].key != "Ahi" && myMemory.buckets[1].key != "Chello" {
		t.Error("Evict is not working as expected")
	}
}

func TestAdd(t *testing.T) {

	myMemory := NewInMemoryStorage(5, 100)
	myMemory.Set("hi", "1212")
	myMemory.Set("hi", "1212")
	myMemory.Set("bye", "1212")
	myMemory.Set("hi", "1212")
	myMemory.Set("hi", "1212")

	myMemory.printBuckets()
	myMemory.printLastUsed()

}
