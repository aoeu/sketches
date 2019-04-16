// This program shows some of the ways that forwarding may be implemented
// in a Go program, while detailing the risks (in the potential for a panic
// to occur when calling the forwarding method) of the easiest approach
// ("unsafeForwarding").
//
// https://en.wikipedia.org/wiki/Forwarding_(object-oriented_programming)

package main

import (
	"fmt"
	"os"
)

type printer interface {
	print() error
}

// unsafeForwarding is a type that features an anonymous field of the "printer" type.
// Because the "printer" field is anonymous, that means the "print()" method of the
// "printer" field is promoted, such that we can call the "print()" method directly from
// a unsafeForwarding instance: "unsafeForwarding.print()"
type unsafeForwarding struct { // this type is the "sender" in a forwarding setup.
	printer // this anonymous field is the "receiver" in a forwarding setup.
}

type safeForwarding struct {
	printer
}

// print is a method prevents the anonymous "printer" field's "print()" method
// from being promoted to the safeForwarding type, and checks if the anonymous
// "printer" field has actually been set before calling its "print()" method
// (i.e. performing the actual "forwarding" of the "print()" method call).
func (s *safeForwarding) print() error {
	if s.printer != nil {
		return s.printer.print()
	}
	return fmt.Errorf("no printer has been set on this 'safeForwarding' struct\n")
}

type statement string

func (s statement) print() error {
	fmt.Println(s)
	return nil
}

type baz struct{}

func (b baz) print() error {
	fmt.Println("baz")
	return nil
}

func main() {
	u := unsafeForwarding{}
	// u.print() // This would pass the compiler, but cause a panic if it executes.
	u.printer = statement("foo")
	if err := u.print(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	s := safeForwarding{}
	if err := s.print(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	s.printer = statement("bar")
	if err := s.print(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// In tests, we could set a mock printer on either forwarding type.
	s.printer = baz{}
	if err := s.print(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
