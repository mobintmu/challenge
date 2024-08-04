package memory

import (
	"fmt"
	"testing"
)

func TestCreateHash(t *testing.T) {
	fmt.Println("hi")
	memory := NewInMemoryStorage(5, 100)
	hash := memory.hash("hi")
	if hash != 1748694682 {
		t.Errorf("Hash function is not working as expected")
	}
}
