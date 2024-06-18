package migrate

import (
	"fmt"

	"github.com/go-kod/kod"
)

type migrate struct {
	kod.Implements[IMigrate]
}

func (migrate) Hello() {
	fmt.Println("hello from migrate")
}
