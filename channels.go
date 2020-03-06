package main

import "fmt"

// ************************************
// channels are the pipes that connects
// concurrent goroutines
// NOTE: by default sends and receives are block
// until both sender and receiver are ready
// *************************************

func main() {

	// create new channel
	// make (chan type)
	msg := make(chan string)

	// send value into a channel using
	// the channel <- syntax
	go func() {
		msg <- "ping"
	}()

	// <- channel syntax receives the value from the channel
	m := <-msg
	fmt.Println(m)
}
