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

func main() {
	fmt.Println(greeting("golang"))
	s, _ := greeting("Santosh")
	fmt.Println(s)
}
