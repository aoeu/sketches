package main 

import (
	"fmt"
)

// Amazon provides golang types that represent an entire API in the following way.
// Note that their interfaces don't end with the idiomatic "er" so neither does this example.
type methods interface {
	foo() string
	bar() string
	baz() string
}

// Let's use 'methods' as an unamed field on a struct, using the interface type itself as a field.
// Because the field is unamed, we'll be able to call the `baz()` function two different ways:
//  d := data{}
//  d.methods.baz()
//  d.baz() // shorthand for `d.methods.baz()`
type data struct {
	methods
}

// And let's explicitly define and bind a 'foo' function to the 'data' struct type.
func (d data) foo() string {
	return "FOO"
}

func main() {
	defer printRuntimeErrors() // Brace for impact, later.
	d := data{}

	// We can call the `foo()` function that was explicitly defined and bound to the `data` struct type.
	fmt.Printf("Result of calling the explictly implemented `data.foo()` function: '%v'\n", d.foo())

	// We can also pass the `data` struct type to a function...
	exec(d)
}

// ...that accepts any type that fulfills the `methods` interface type,
// which the `data` struct does, but only *implictly* via the unnamed `methods` field on the `data` struct definition,
// not because we've truly fulfilled the interface by defining and binding `foo()`, `bar()`, and `baz()` functions
// to the `data` struct type!
func exec(m methods) {

	// This means we can call the `foo()` function of whatever was sent as a parameter,
	// and what will really be called is the explicitly defined `foo()` function we definined on the `data` struct type.
	fmt.Printf("Result of calling the explictly implemented `data.foo()` function: '%v\n'", m.foo())

	// We can also try to call the `bar()` function that was implictly defined by the unnamed `methods` field 
	// on the `data` struct definiton.
	fmt.Println(m.bar())
}

func printRuntimeErrors() {
	// The result was a segmentation error because a `bar()` method was never explictly defined and bound
	// to the `data` struct type. We've kept the compiler from stopping us by using the `method` interface
	// on the `data` struct type as a field itself, instead of manually creating and binding functions
	// to the `data` struct type.
	if r := recover(); r != nil {
		fmt.Println("Tried to call `m.bar()` and crashed:", r)
	}

	// In other words, explictly defined functions of interface methods, bound to a struct type,
	// enables the *language* and compiler determine whether or not the interface is fulfilled.

	// But explicitly defining the struct and including the interface definition itself as a field,
	// implictly means the interface methods are defined and will casue a panic if they are called!
}
