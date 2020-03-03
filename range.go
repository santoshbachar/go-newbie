package main

import "fmt"

func main() {
	// *******************************************
	// range works on a variety of data structures
	// *******************************************

	// on arrays / slices
	array := [5]int{1, 2, 3, 4, 5}
	for key, val := range array {
		fmt.Println("key: ", key, " value: ", val)
	}

	// key can be ommited using _
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Multipliers of 2 are: ")
	for _, n := range nums {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}

	// range requires both key,val pairs in maps
	chart := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for k, v := range chart {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range with keys only
	for key := range chart {
		fmt.Println("key: ", key)
	}

	// range on string
	// itirates over Unicode code points
	for k, v := range "go" {
		fmt.Println(k, v)
	}

}
