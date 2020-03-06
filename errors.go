package main

import (
	"errors"
	"fmt"
)

func greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Name cannot be empty")
	}
	return "hello " + name, nil
}

// custom type as error
type myError struct {
	arg  int
	prob string
}

// implement Error() on type myError
// to make custom error
func (e myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func calc(a int) (int, error) {
	if a < 0 {
		return -1, myError{a, "negative number is invaid"}
	}
	return a * a, nil
}

func main() {
	fmt.Println(greeting("golang"))
	s, _ := greeting("Santosh")
	fmt.Println(s)

	for _, v := range []int{2, -3, -4} {
		if val, err := calc(v); err != nil {
			fmt.Println("calc error: ", err)
		} else {
			fmt.Println("calc: ", val)
		}
	}

	// get the error as an instance of the custom error
	// via type assertion
	_, e := calc(-1)
	if ae, ok := e.(myError); ok {
		fmt.Println("arg: ", ae.arg)
		fmt.Println("prob: ", ae.prob)
	}
}
