package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	t, err := time.ParseInLocation("2006-01-02T15:04:05.000", "2019-09-16T09:09:30.000", newYork)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.String())
}
