package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Increment i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// for init
	for j := 0; j < 10; j++ {
		fmt.Println("Increment j", j)
		time.Sleep(time.Second)
	}

	// fmt.Println(j) its not possible use variable j out of loop scope

	names := [3]string{"foo", "bar", "barz"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// 0 foo
	// 1 bar
	// 2 barz

	for _, name := range names {
		fmt.Println(name)
	}

	// foo
	// bar
	// barz

	for index, letter := range "WORD" {
		fmt.Println(index, letter)
	}

	// 0 87
	// 1 79
	// 2 82
	// 3 68

	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	// 0 W
	// 1 O
	// 2 R
	// 3 D

	user := map[string]string{
		"first_name": "Peter",
		"last_name":  "Parker",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	// first_name Peter
	// last_name Parker

	// Infinite loop syntax

	/*
		for {
			...
		}
	*/

}
