package main

import "fmt"

// ***********************************
// struct = typed collection of fields
// ***********************************

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 20
	return &p
}

func main() {
	fmt.Println(person{"Santosh", 25})
	fmt.Println(person{name: "Katty Perry", age: 50})
	fmt.Println(person{name: "Narendra Modi"})
	fmt.Println(&person{name: "Larry Page", age: 45})
	fmt.Println(NewPerson("Sergey Brin"))

	s := person{name: "Mark", age: 39}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 35
	fmt.Println(s.age)
}
