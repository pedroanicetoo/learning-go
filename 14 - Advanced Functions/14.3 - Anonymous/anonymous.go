package main

import "fmt"

func main() {
	func(text string) {
		fmt.Println(text)
	}("Hello World") // after declare anonymous function, for execute it, you have to pass '()' declaration

	// another example

	return1 := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Hello World")

	fmt.Println(return1)
}
