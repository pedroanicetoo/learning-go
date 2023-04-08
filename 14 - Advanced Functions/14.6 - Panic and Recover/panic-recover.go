package main

import "fmt"

/*
 ``Panics fall into the second category of errors, which are unanticipated by the programmer.
 These unforeseen errors lead a program to spontaneously terminate and exit the running Go program.
 Common mistakes are often responsible for creating panics.
 There are certain operations in Go that automatically return panics and stop the program.
 Common operations include indexing an array beyond its capacity, performing type assertions, calling methods on nil pointers, incorrectly using mutexes,
 and attempting to work with closed channels.
 Most of these situations result from mistakes made while programming that the compiler has no ability to detect while compiling your program.
 Since panics include detail that is useful for resolving an issue,
 developers commonly use panics as an indication that they have made a mistake during a program’s development.``
*/

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Recovery execution with success!")
	}
}

func studentApproved(n1, n2 float64) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is EXACTLY 6!")
}

func main() {
	fmt.Println(studentApproved(6, 7))
	fmt.Println("After execution!")
	// true
	// After execution!
	fmt.Println(studentApproved(6, 6))
	fmt.Println("After execution!")
	// Recovery execution with success!
	// false
	// After execution!
}
