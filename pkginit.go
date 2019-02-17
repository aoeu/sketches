package main

import "fmt"

var (
	foo = "foo"
	bar = "bar"
)

// init is called after all the variable declarations in the package have
// evaluated their initializers, and those are evaluated only after all
// the imported packages have been initialized. Each source file may have
// an init function.
//
// https://golang.org/doc/effective_go.html#init
//
func init() {
	foo = "baz"
	bar = "qux"
}

func main() {
	fmt.Println(foo, bar) // prints "baz qux"
}