package main

import "fmt"

func main() {
	// Copy memory address a in b
	a := 2
	b := &a
	fmt.Println(b, &a)
}
