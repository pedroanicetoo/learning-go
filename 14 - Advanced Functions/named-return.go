package main

import "fmt"

func mathCalculations(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return // don't need to declare which variable should return because assignments return already did that
}

func main() {
	sum, sub := mathCalculations(10, 20)
	fmt.Println(sum, sub)
}
