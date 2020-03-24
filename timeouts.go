package main

import (
	"fmt"
	"time"
)

// **********************************************************************
// timeouts are important for programs that connect to external resources
// or bound to time execution
// NOTE: Go implement timeout using chanels and select
// **********************************************************************

func main() {

	// simulating an external call, that returns its result on a channel after 2 seconds
	// NOTE: channel is buffered, so the send in the goroutine is non-blocking.
	// This apporoach prevents go routine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// the select awaits and proceeds with <-time.After as it arrives 1 second earlier than <-c1
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
