package main

import "fmt"

func main() {
	// aritmetics operators

	var n1 int16 = 1
	var n2 int16 = 1
	n3 := n1 + n2 // only works if both variables are the same type

	fmt.Println(n3)

	// assignments operators

	var var1 string = "string 1"
	var2 := "string 2"
	fmt.Println(var1, var2)

	// relation operators

	fmt.Println(1 > 2, 1 < 2, 1 == 2, 1 != 2)

	// logic operators

	fmt.Println(true && false, true || false, !true)

	// unary operators

	n := 1
	n++
	n += 10
	fmt.Println(n)

	// ternary operator

	var text string

	if n > 10 {
		text = "bigger than 10"
	} else {
		text = "less or equal than 10"
	}
	fmt.Println(text)
}
