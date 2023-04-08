package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, numbers ...int) { // can have only one parameter variatic, and MUST be the last
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	total_sum := sum(1, 2, 3, 4, 5)
	fmt.Println(total_sum)

	write("Hello World", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
