package dump

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func channels() {
	// messages := make(chan string)
	// go func() { messages <- "ping" }()
	// msg := <-messages
	// fmt.Println(msg)

	go channelBuffering("! ")
	go channelBuffering("* ")

	done := make(chan bool, 1)
	go channelSync(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func channelBuffering(idx string) {
	max := 5
	ch := make(chan string, max)

	for i := 0; i < max; i++ {
		ch <- idx + strconv.Itoa(rand.Intn(max))
	}
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

func channelSync(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
