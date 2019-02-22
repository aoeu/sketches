package main

import (
	"fmt"
	"log"
	"math/rand"
)

type foo struct{}

func main() {
	n := rand.Intn(50)
	c := make(chan *foo, n) // We'll use a pointer type to make life harder / make pushing nil into the channel possible.
	v := &foo{}

	for i := 1; i < n; i++ { // 1) Change '1' to '0' in this line after addressing comment #2 to make the test pass.
		c <- v
	}
	c <- nil // 2) Remove this line to observe a different test failure case, then address comment #1.
	close(c)

	if err := testChannelIsFullOfNonNilValues(c); err != nil {
		log.Fatal("failed test: ", err)
	}
	fmt.Println("passed test")

}

func testChannelIsFullOfNonNilValues(c chan *foo) error {
	n := len(c)
	if n != cap(c) {
		return fmt.Errorf("expected channel of %v values but there were only %v", cap(c), n)
	}

	i := 0
	for v := range c {
		i++
		if v == nil {
			return fmt.Errorf("expected value in channel but received nil")
		}
	}

	if i != cap(c) {
		return fmt.Errorf("only read %v values before channel closed, but expected %v values", i, cap(c))
	}
	return nil
}
