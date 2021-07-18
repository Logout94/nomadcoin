package main

import "fmt"

func main() {
	// array have a finite
	foods := [3]string{"potato", "pizza", "pasta"}

	for _, dish := range foods {
		fmt.Println(dish)
	}
}
