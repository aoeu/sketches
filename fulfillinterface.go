package main

import (
	"fmt"
)

type substance string

const (
	slime substance = "slime"
	ink   substance = "ink"
)

type oozer interface {
	ooze() (out substance)
}

type octopus struct {
	substance
}

func newOctopus() *octopus {
	return &octopus{substance: ink}
}

func (o *octopus) ooze() substance {
	return o.substance
}

type slug struct {
	substance
}

func newSlug() *slug {
	return &slug{substance: slime}
}

func (s *slug) ooze() substance {
	return s.substance
}

func main() {
	fmt.Println(newOctopus().ooze())
	fmt.Println(newSlug().ooze())
}
