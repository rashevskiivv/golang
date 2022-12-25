package dump

import "fmt"

// non-blocking operations
func chanOper() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	func() {
		messages <- "1"
		signals <- true
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	case sig := <-signals:
		fmt.Println("signal received", sig)
	default:
		fmt.Println("no activity")
	}
}
