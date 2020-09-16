// https://play.golang.org/p/G-QTPSvP7-1
package main

import (
	"fmt"
)

type (
	key   interface{}
	value interface{}
)

var m map[key]value

func init() {
	m = make(map[key]value)
}

func main() {
	m["foo"] = "bar"
	m[1] = "baz"
	m['z'] = 0xFF
	fmt.Println(m)
}
