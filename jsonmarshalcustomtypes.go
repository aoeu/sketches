package main

// https://play.golang.org/p/I8bOTi6oReE

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type birthdate time.Time

func (b birthdate) String() string {
	y, m, d := time.Time(b).Date()
	return fmt.Sprintf("%d-%d-%d", y, m, d)
}

func (b birthdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DateOfBirth string `json:"dob"`
	}{
		DateOfBirth: b.String(),
	})
}

func main() {
	d := birthdate(time.Now())
	fmt.Printf("The String() method is called for both string and as opaque value Printf args: '%s' == '%v'\n", d, d)
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not marshal custom type to JSON due to error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("JSON payload is: %v\n", string(b))
}
