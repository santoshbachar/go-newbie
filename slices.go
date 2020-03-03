package main

import (
	"fmt"
	"sort"
)

func main() {

	// create empty slice with non-zero length
	s := make([]string, 3)
	fmt.Println("slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[1])

	fmt.Println("length of slice: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice: ", c)

	// simple way to copy a slice
	cc := s
	fmt.Println("cc = ", cc)

	// slice operator
	// slice[low:high]
	l := s[2:5]
	fmt.Println("s[2:5] = ", l)

	fmt.Println("s[:5] is ", s[:5])
	fmt.Println("s[2:] is ", s[2:])

	// single line slice declaration and initialization
	ss := []string{"I", "Love", "Go"}
	fmt.Println("ss: ", ss)

	// slice with capacity
	sc := make([]int, 3, 9)
	fmt.Println(sc)

	// iterating slice with for loop
	for i := 0; i < len(sc); i++ {
		sc[i] = i + 1
	}
	fmt.Println(sc)
	fmt.Println("sc length: ", len(sc))
	fmt.Println("sc capacity: ", cap(sc))

	ns := []int{2, 3, 5, 6, 7, 4, 1, 9, 10, 8}
	fmt.Println("ns before sort: ", ns)

	// requires import "sort"
	sort.Ints(ns)
	fmt.Println("ns after sort: ", ns)

	// works with slices of string type as well
	nsi := []string{"New", "Delhi", "is", "the", "capital", "of", "India"}
	fmt.Println("nsi before sort: ", nsi)
	sort.Strings(nsi)
	fmt.Println("nsi after sort: ", nsi)

}
