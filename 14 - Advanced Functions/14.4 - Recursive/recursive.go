package main

import "fmt"

func fibonnaci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonnaci(position-2) + fibonnaci(position-1)
}

// 1 1 2 3 5 8 13 21

func main() {
	position := uint(12)

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonnaci(i))
	}
}
