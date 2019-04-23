package main

import "fmt"

func main() {
	// The "and not" operator ( "&^") functions somewhat like a xor gate,
	// but only accepting high bits from the left operand.
	// For example, if 8-bit representations the decimal numbers 3 and 1 were
	// "and notted" together, the result would be decimal number 2:
	// decimal 3 as uint8 binary:  00000011
	// decimal 1 as uint8 binary:  00000001
	// the "and not" result:       00000010
	var three, one uint8 = 3, 1
	two := three &^ one
	fmt.Println(two)
}