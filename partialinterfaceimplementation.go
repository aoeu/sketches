package main

import (
	"fmt"
	"os"
	"strings"
)

type shouter interface {
	shout(string) string
}

type alice struct {
}

type bob struct {
}

func (b bob) shout(s string) string {
	return strings.ToUpper(s)
}

type chuck struct {
	shouter
}

func shoutAtStdout(s shouter, phrase string) {
	fmt.Fprintln(os.Stdout, s.shout(phrase))
}

func main() {
	// The following line would not pass the compiler if uncommented and recompiled.
	// shoutAtStdout(new(alice), "I stop compilation because the 'alice' type has no 'shout' method!")

	// The following line does probably what you'd expect.
	shoutAtStdout(new(bob), "I fulfilled the shouter interface normally by defining and binding a function to bob!")

	// The following passes the compiler, leads to a panic due to a 'nil pointer dereference' error at runtime.
	shoutAtStdout(new(chuck), "I tricked the compiler using the shouter interface type itself as an embedded field!")
}
