package main

import "fmt"

var n int

func main() {
	fmt.Println("Function main has been executed")
	fmt.Println(n)
}

func init() {
	fmt.Println("Function init has been executed")
	n = 10
}

// Function init has been executed
// Function main has been executed
// 10
