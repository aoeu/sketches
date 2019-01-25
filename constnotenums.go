package main

import "fmt"

type suit rune

const (
	spades   = suit('♠')
	hearts   = suit('♥')
	diamonds = suit('♦')
	clubs    = suit('♣')
	empty    = suit('_') // Joker
	unknown  = suit('?') // Cups, Bells, Acorns, Sheilds, etc.
)

func (s suit) name() string {
	// Switch on the value of suit type just like switching on values of an enum class in languages like "yava."
	switch s {
	case spades:
		return "Spades"
	case hearts:
		return "Hearts"
	case diamonds:
		return "Diamonds"
	case clubs:
		return "Clubs"
	case empty:
		return "Empty"
	case unknown:
		return "Unknown"
	default:
		return "Undefined"
	}
}

// String concatenates the symbol and name of the suit as a string.
func (s suit) String() string {
	return string(s) + " : " + s.name()
}

func main() {
	fmt.Printf("Card suits:\n %v\n %v\n %v\n %v\n", spades, hearts, diamonds, clubs)
}
