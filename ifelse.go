package main

import "fmt"

func main() {

	// ***************************
	// there's no ternary if in go
	// ***************************

	earthIsFlat := false

	// basic if/else
	if earthIsFlat {
		fmt.Println("You're crazy AF")
	} else {
		fmt.Println("Hmm... Good, you know the facts right !")
	}

	// if without else
	if 2020%4 == 0 {
		fmt.Println("2020 is a leap year")
	}

	// declare variable, then branch
	// also mind 'else if'
	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}

}
