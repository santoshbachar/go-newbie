package main

import "fmt"

func main() {
	num := make(map[string]int)

	num["one"] = 1
	num["two"] = 2
	num["three"] = 3

	fmt.Println(num)

	t := num["two"]
	fmt.Println(t)

	fmt.Println("length of num: ", len(num))

	delete(num, "two")
	fmt.Println("num after delete: ", num)

	// check if key is present
	_, present := num["two"]
	fmt.Println(present)

	// declare and initialize in same line
	n := map[string]string{"India": "New Delhi", "Russia": "Moscow"}
	fmt.Println(n)
}
