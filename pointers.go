package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(val *int) {
	*val = 0
}

func main() {
	a := 2
	pa := &a

	fmt.Println("a=", a)
	fmt.Println("pa=", pa)

	i := 1
	fmt.Println("Initial i value: ", i)

	zeroval(i)
	fmt.Println("after zeroval: ", i)

	// & gives address of i
	zeroptr(&i)
	fmt.Println("after zeroptr: ", i)
}
