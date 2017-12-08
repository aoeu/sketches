package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var RNG = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	maximumMillisecondsToWait = 3000
	totalNumberOfTasks        = 10
)

type task func(wg *sync.WaitGroup)

func makeTasks(n int) []task {
	t := make([]task, n)
	for i := 0; i < n; i++ {
		s := time.Duration(RNG.Intn(maximumMillisecondsToWait)) * time.Millisecond
		out := fmt.Sprintf("Task %v completed after %v", i, s)
		t[i] = func(wg *sync.WaitGroup) {
			defer wg.Done()
			<-time.After(s)
			fmt.Println(out)
		}
	}
	return t
}

func main() {
	var wg sync.WaitGroup
	for _, task := range makeTasks(totalNumberOfTasks) {
		wg.Add(2)
		go task(&wg) // This does what you might expect...
		go func() {
			task(&wg) // ... and this might not. (The closure is compiled once! h/t @wizardishungry)
		}()
	}
	wg.Wait()
}
