package store

import (
	"fmt"

	"github.com/go-kod/kod"
)

type DiskStore struct {
	kod.Implements[IKodComponent]
}

func (DiskStore) Get(string) any {
	fmt.Println("dummy get called")
	return "dummy string"
}

func (DiskStore) Set(string, any) error {
	fmt.Println("dummy set called")
	return nil
}
