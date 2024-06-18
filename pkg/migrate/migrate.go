package migrate

import (
	"fmt"

	"github.com/go-kod/kod"
)

type Migrate struct {
	kod.Implements[IMigrate]
}

func (Migrate) Hello() {
	fmt.Println("hello from migrate")
}
