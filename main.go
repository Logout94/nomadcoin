package main

import "fmt"

func plus(a ...int) int {
	var total int
	for _, item := range a { // _ = ignore
		total += item
	}
	return total
}

func main() {
	result := plus(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}
