package main

import "fmt"

func function1() {
	fmt.Println("Execute function1")
}

func function2() {
	fmt.Println("Execute function2")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("Average calculated. The Result will be returned.") // In this example, this line will be executed before of function return
	fmt.Println("Entry in the studentApproved function")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	// defer function1() // defer this function until last possible moment before the return
	// function2()
	fmt.Println(studentApproved(7, 8))

}
