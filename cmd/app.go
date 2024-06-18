package cmd

import (
	"github.com/go-kod/kod"
)

type app struct {
	kod.Implements[kod.Main]
	demo1 kod.Ref[Demo1]
	demo2 kod.Ref[Demo2]
}
