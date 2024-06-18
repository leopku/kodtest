package store

import (
	"fmt"

	"github.com/go-kod/kod"
)

type MemoryStore struct {
	kod.Implements[IStore]
}

func (MemoryStore) Get(string) any {
	fmt.Println("dummy get called")
	return "dummy string"
}

func (MemoryStore) Set(string, any) error {
	fmt.Println("dummy set called")
	return nil
}
