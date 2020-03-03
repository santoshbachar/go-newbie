package main

import (
	"fmt"
	"math"
)

const capital string = "Delhi"

func main() {
	fmt.Println(capital)

	const population = 2500000

	// constant expressions perform
	// arthematic operations with arbitrary precision.
	const grainsInStock = 3e20 / population

	fmt.Println(grainsInStock)

	// numeric constant has no type
	// until explicit conversion.
	fmt.Println(int64(grainsInStock))

	// math requires import "math"
	fmt.Println(math.Sin(2))
}
