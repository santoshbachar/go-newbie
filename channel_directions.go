package main

import "fmt"

// *********************************************
// when using channels in function parameters,
// you can specify if a channel is meant to only
// send or receive values.
// *********************************************

// ping function only accept a channel for sending values.
// this increases type-safety
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong function accepts one channel for receives (pings)
// and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
