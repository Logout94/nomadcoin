package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println("Main 'nico':", nico)
}
