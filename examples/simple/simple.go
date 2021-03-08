package main

import (
	"fmt"
)

type ErrNo int

const (
	ErrFoo ErrNo = iota // Foo happened, this cannot be
	ErrBar              // Bar occurred, whoops
	ErrBaz              // Baz Baz Baz
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(ErrFoo.Error())
}
