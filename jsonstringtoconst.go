// This program demonstrates how a JSON string can be unmarshalled as a constant value of a type that derives from the string type.
// This approach may be used to mimic unboxing a JSON string into a predefined enum value (which is just a constant in Go).
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Bar string

type Payload struct {
	Foo Bar `json:"baz"`
}

var s = `
{
	"baz" : "qux"
}
`

const qux = Bar("qux")

func main() {
	var p Payload
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("field unmarshalled from a JSON string to the 'Bar' type, in field named 'Foo': %+v\n", p)
	if p.Foo == qux {
		fmt.Println("unmarshalled JSON string value is equal to the constant value of a non string type (i.e. an enum value)")
	}
}
