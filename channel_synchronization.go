package main

import (
	"fmt"
	"time"
)

// *************************************************************
// channels are used to synchronize execution across goroutines.
// blocing receive can be used to wait for a goroutine to finish.
// *************************************************************

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
