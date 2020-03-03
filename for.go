package main

import "fmt"

func main() {

	// *******************************
	// for is Go's only loop construct
	// *******************************

	// single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition/after for loop
	for j := 7; j > 1; j-- {
		fmt.Println(j)
	}

	// without condition
	// loop until break
	for {
		fmt.Println("loop till break")
		break
	}

	// next itiration with 'continue'
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
