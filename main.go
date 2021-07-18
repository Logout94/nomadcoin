package main

import "fmt"

func main() {
	// array have a finite
	foods := [3]string{"potato", "pizza", "pasta"}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
}
