package main

import "fmt"

func swapSignal(number int) int {
	return number * -1
}

func swapSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	upsidedownNumber := swapSignal(number)
	fmt.Println(upsidedownNumber) // -20
	fmt.Println(20)               // 20 (does not change value from memory address)

	// using pointer...
	new_number := 40
	swapSignalWithPointer(&new_number)
	fmt.Println(new_number) // -40
}
