package main

import (
	"challenge/memory"
)

func main() {

	memory := memory.NewInMemoryStorage(5, 100)

	_ = memory
}
