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
		key:       "hi",
		value:     "1212",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	B := &data{
		key:       "bye",
		value:     "1312",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	C := &data{
		key:       "hello",
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

	myMemory.printLastUsed()

	if myMemory.lastUsed[0].key != "hi" && myMemory.lastUsed[1].key != "hello" && myMemory.lastUsed[2].key != "bye" {
		t.Errorf("LRU is not working as expected")
	}

}
