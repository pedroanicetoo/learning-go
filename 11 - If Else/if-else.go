package main

import "fmt"

func main() {
	fmt.Println("Control Structures:")

	number := 0

	if number > 15 {
		fmt.Println("bigger than 15")
	} else {
		fmt.Println("less or equal than 15")
	}

	// if init
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("number is bigger than 0")
	} else if number == 0 {
		fmt.Println("number is 0")
	} else {
		fmt.Println("number is less than 0")
	}

	// anotherNumber variable is available to use only at if init scope, if it's used later it will be a unknowing variable

}
