package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// Note: only final parameter have type
// it implies for the rest
func add3(a, b, c int) int {
	return a + b + c
}

func main() {
	add := sum(1, 2)
	fmt.Println("1+2 = ", add)

	fmt.Println("1+2+3 = ", add3(1, 2, 3))
}
