package main

import "fmt"

func main() {
	// Basic scenario
	fmt.Println("Pointers")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2) // 10 10

	var1++
	fmt.Println(var1, var2) // 11 10

	// Pointer is a memory reference

	var var3 int
	var pointer *int

	var3 = 100
	pointer = &var3

	fmt.Println(var3, pointer) // 100 0xc0000180f8
	var3++
	fmt.Println(var3, *pointer) // 100 100 (using dereferencing)
}
