package main

import "fmt"

func add(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalculation(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	fmt.Println(add(2, 3))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("function text")
	fmt.Println(result)

	result_sum, result_sub := mathCalculation(1, 2)
	result_sum2, _ := mathCalculation(1, 2) // ignoring second return
	fmt.Println(result_sum, result_sub)
	fmt.Println(result_sum2)
}
