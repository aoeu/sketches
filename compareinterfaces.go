package main

import (
	"fmt"
)

type namer interface {
	name() string
}

type backwardsNamer struct {
	origName string
}

func (b backwardsNamer) name() string {
	s := ""
	for i := len(b.origName) - 1; i >= 0; i-- {
		s += string(b.origName[i])
	}
	return s
}

type foo struct {
	namer *namer
}

func main() {
	n := namer(backwardsNamer{"travis"})
	b := foo{&n}
	m := namer(backwardsNamer{"travis"})
	f := foo{&m}
	if *b.namer != *f.namer {
		fmt.Println("different values")
	}
	if b.namer != f.namer {
		fmt.Println("different pointer values")
	}
}
