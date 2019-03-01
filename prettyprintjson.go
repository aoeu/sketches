package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func prettyPrint(payload []byte) {
	b := new(bytes.Buffer)
	if err := json.Indent(b, payload, "", "    "); err != nil {
		fmt.Printf("could not debug print '%v' due to error: %v", string(payload), err)
	} else {
		fmt.Printf("%s", b)
	}
}

func main() {
	prettyPrint([]byte(`{ "foo":"bar", "baz" :1}`))
}
