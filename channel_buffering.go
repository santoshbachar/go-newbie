package main

import "fmt"

// ***************************************************
// by default channels are unbuffered
// i.e., they will only accept sends (chan <-)
// only if there is receive (<- chan) ready to
// receive the send value.
// Buffered channels accept a limited number of values
// without a corresponding receiver for those values
// ***************************************************

func main() {

	// channel of strings buffering upto 2 values
	msg := make(chan string, 2)

	msg <- "India"
	msg <- "Israel"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
