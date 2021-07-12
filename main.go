package main

import "fmt"

func main() {
	name := "Nicolas!!!!!!! Is my name"
	for _, letter := range name {
		//fmt.Println(letter)
		fmt.Printf("%b\n", letter)
	}
}
