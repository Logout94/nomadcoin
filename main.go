package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello My name is %s and I'm %d\n", p.name, p.age)
}

func (p person) KoreanAge() {
	fmt.Printf("My Korean age is %d.\n", p.age)
}
func main() {
	nico := person{name: "nico", age: 12}
	nico.sayHello()
	nico.KoreanAge()
}
