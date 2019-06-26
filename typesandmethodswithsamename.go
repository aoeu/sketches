package main

import (
	"fmt"
)

type foo string

type bar string

func (b bar) foo() foo {
	return foo(b)
}

func main() {
	fmt.Println(foo("Hello,"), bar("world").foo())
}
