// This program demonstrates how nested JSON objects can be unmarshalled into anonymous fields of struct types.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Payload struct {
	Zip `json:"zip"` // Note that the annotations are required for this to work.
}

type Zip struct {
	Ding `json:"ding"`
}

type Ding struct {
	Pop `json:"pop"`
}

type Pop string

var s = `
{
	"zip" : {
		"ding" : {
			"pop" : "poof"
		}
	}
}
`

func main() {
	var p Payload
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("nested JSON objects unmarshalled correctly into anonymous fields: %+v\n", p)

}
