package main

import (
	"log"
	"time"
)

func wait(throttle chan bool, t time.Duration, id int) {
	log.Printf("goroutine %v : enter\n", id)
	<-time.After(t)
	throttle <- true
	log.Printf("goroutine %v : exit\n", id)
}

func makeThrottle(maxGoRoutines int) chan bool {
	throttle := make(chan bool, maxGoRoutines)
	for i := 0; i < maxGoRoutines; i++ {
		throttle <- true
	}
	return throttle
}

func main() {
	maxGoRoutines := 5
	t := makeThrottle(maxGoRoutines)
	for i := 0; i < 20; i++ {
		<-t
		go wait(t, 1*time.Second, i)
	}
	log.Println("Done")
}
