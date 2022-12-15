package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var completed = false

func onceFunc() {
	var wg sync.WaitGroup
	var once sync.Once
	loopConst := 100
	wg.Add(loopConst)

	for i := 0; i < loopConst; i++ {
		go func() {
			if found() {
				once.Do(markCompleted)
				// markCompleted()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func markCompleted() {
	completed = true
	fmt.Println("completed")
}

func found() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
