package main

import (
	"fmt"
	"time"
)

// ****************************************************
// select lets you wait on multiple channel operations.
// ****************************************************

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// simulating blocking RPC operation, executing
	// in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("c1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("c2 received", msg2)
		}
	}

}
