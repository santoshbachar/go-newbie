package main

import "fmt"

// ******************************************************
// Sends and receives on channels are blocking by deafult
// select with a default clause can be used to implement
// non-blocking sends, receives and even
// non-blocking multi-way selects
// ******************************************************

func main() {
	messages := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
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
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("signal received", sig)
	default:
		fmt.Println("no activity")
	}
}
