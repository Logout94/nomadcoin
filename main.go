package main

import "fmt"

func main() {
	x := 49495495945949
	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	xAsBin := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBin)
}
