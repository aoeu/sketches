package main

import (
	"fmt"
)

type name string

type first name
type middle name
type last name

type fullName struct {
	first
	middle
	last
}

func (f fullName) String() string {
	return fmt.Sprintf("%s %s %s", f.first, f.middle, f.last)
}

func main() {
	n := "fido"
	if n == "fido" {
		fmt.Println("Type inference allows us to still type a 'name' value as if we're assigning to a string variable.")
	}

	// var s string = "fido"
	// if n == s {
	//	fmt.Println("This block can't build due to: 'invalid operation: n == m (mismatched types name and string)"'")
	// }

	b := fullName{first: "Billy", middle: "Joe", last: "Armstrong"}
	j := fullName{first: "Joe", last: "Strummer"}

	if name(b.middle) == name(j.first) {
		fmt.Println("Joe is a common first name or middle name.")
	}

	// if b.middle == j.first {
	//	fmt.Println("This block can't build due to: 'invalid operation: b.middle == j.first (mismatched types middle and first)"')
	// }
}
