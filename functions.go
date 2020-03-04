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

func return2Val() (int, int) {
	return 2, 3
}

// variadic function
// n number of arguments
func variadicFunc(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	add := sum(1, 2)
	fmt.Println("1+2 = ", add)

	fmt.Println("1+2+3 = ", add3(1, 2, 3))

	a, b := return2Val()
	fmt.Println("return2Val() gives ", a, b)

	// subset of return value
	_, c := return2Val()
	fmt.Println("only the second value is ", c)

	variadicFunc(1, 2)
	variadicFunc(1, 2, 3)
	variadicFunc(1, 2, 3, 4)
}
