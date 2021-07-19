package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello My name is %s and I'm %d\n", p.name, p.age)
}
func main() {
	nico := person{name: "nico", age: 12}
	nico.sayHello()
}
