package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	args := struct {
		from string
		to   string
	}{
		from: "2019-09-20",
		to:   "2019-10-01",
	}
	layout := "2006-01-02"
	from, err := time.Parse(layout, args.from)
	if err != nil {
		log.Fatal(err)
	}
	to, err := time.Parse(layout, args.to)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= int(to.Sub(from).Hours()/24); i++ {
		fmt.Println(from.AddDate(0, 0, i).Format(layout))
	}
}
