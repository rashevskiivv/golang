package main

import "fmt"

func rangeChan() {
	// queue := make(chan string, 2)
	// queue <- "one"
	// queue <- "two"
	queue := make(chan int, 2)
	queue <- 1
	queue <- 2
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
