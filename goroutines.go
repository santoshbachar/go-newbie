package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// run synchronously
	f("direct")

	// run concurrently
	go f("goroutine")

	// goroutine for anonymous function call
	go func(msg string) {
		fmt.Println("msg: ", msg)
	}("going")

	// waiting for goroutines to finish
	// time requires import "time"
	time.Sleep(time.Second)
	fmt.Println("done")
}
