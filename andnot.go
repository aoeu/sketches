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

	// The operator could arguably be called "left xor" instead of "and-not"
	// in order to avoid any conflation with "nand" (not-and).

	// This can be a little unexpected, especially because "|^" is not an
	// operator - the one would be bitwise negated (bit-flipped), then
	// "or" gated with the number three, resulting in 255 (all 8 bits being high-bits).
	twoHundredFiftyFive := three | ^one
	fmt.Println(twoHundredFiftyFive)

	// Note how the formatter forcibly justifies the negate to the right operand.
	// That help go to how having a universal formatter in a language can help
	// even in areas where the language syntax can be tricky.
}