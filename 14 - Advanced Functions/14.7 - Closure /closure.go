package main

import "fmt"

// functions that reference variables out of its scope

func closure() func() {
	text := "inside of closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "inside of main function"
	fmt.Println((text))

	new_function := closure()

	new_function()
}
