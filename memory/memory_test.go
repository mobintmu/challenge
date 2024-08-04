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
		key:       "A-hi",
		value:     "1212",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	B := &data{
		key:       "B-bye",
		value:     "1312",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	C := &data{
		key:       "C-hello",
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

	if myMemory.lastUsed[0].key != "A-hi" && myMemory.lastUsed[1].key != "C-hello" && myMemory.lastUsed[2].key != "B-bye" {
		t.Errorf("LRU is not working as expected")
	}

}

func TestRemoveEvict(t *testing.T) {

	A := &data{
		key:       "A-hi",
		value:     "1212",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	B := &data{
		key:       "B-bye",
		value:     "1312",
		timestamp: time.Now().Add(1 * time.Hour).Unix(),
	}

	C := &data{
		key:       "C-hello",
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

	if myMemory.buckets[0].key != "A-hi" && myMemory.buckets[1].key != "C-hello" {
		t.Error("Evict is not working as expected")
	}
}

func TestAdd(t *testing.T) {

	myMemory := NewInMemoryStorage(5, 100)
	myMemory.Set("hi", "1212", time.Now().Unix()+3600)
	myMemory.Set("hi", "1212", time.Now().Unix()+3600)
	myMemory.Set("bye", "1212", time.Now().Unix()+3600)
	myMemory.Set("hi", "1212", time.Now().Unix()+3600)
	myMemory.Set("hi", "1212", time.Now().Unix()+3600)

	// myMemory.printBuckets()
	// myMemory.printLastUsed()

	if myMemory.buckets[0].key != "hi" && myMemory.buckets[1].key != "bye" {
		t.Error("Add is not working as expected")
	}

}

func TestSetAndGetWithTime(t *testing.T) {

	myMemory := NewInMemoryStorage(5, 100)
	myMemory.Set("hi", "1212", time.Second.Nanoseconds())
	_, value_bool := myMemory.Get("hi")
	if value_bool == false {
		t.Errorf("Get is not working as expected")
	}
	time.Sleep(2 * time.Second)

	_, value_bool = myMemory.Get("hi")
	if value_bool == true {
		t.Errorf("Get is not working as expected")
	}

}

func TestGetAndSetTime(t *testing.T) {
	myMemory := NewInMemoryStorage(3, 100)
	myMemory.Set("hi", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("bye", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("hello", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("hi", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("hi", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("bye", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)
	myMemory.Set("no", "1212", time.Second.Nanoseconds())
	// fmt.Println(myMemory.size)

	// myMemory.printBuckets()
	// myMemory.printLastUsed()

	_, value_bool := myMemory.Get("hello")
	if value_bool == true {
		t.Errorf("Get is not working as expected")
	}

}
