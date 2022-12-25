package dump

import (
	"fmt"
)

func selectFunc() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "one"
	}()

	go func() {
		c2 <- "two"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("default case")
		}
	}
}
