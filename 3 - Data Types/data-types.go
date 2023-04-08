package main

import (
	"errors"
	"fmt"
)

func main() {
	// integer types
	var num1 int8 = 10
	// uint8 alias
	var num2 byte = 10
	var num3 int16 = 10
	var num4 int32 = 10
	// int32 alias
	var num5 rune = 10
	var num6 int64 = 10
	var num7 uint = 10 // unsigned
	num8 := 10

	fmt.Println(num1, num2, num3, num4, num5, num6, num7, num8)

	// float types
	var real1 float32 = 123.45
	var real2 float64 = 123000000000.45
	real3 := 12345.67

	fmt.Println(real1, real2, real3)

	// string and char
	var str1 string = "text"
	str2 := "text2"
	char := 'B' // uint32/rune type

	fmt.Println(str1 + "\n" + str2)
	fmt.Println(char)

	// default value
	var text string
	fmt.Println(text) // ""
	var text2 int16
	fmt.Println(text2) // 0

	// boolean and error

	var boolean1 bool = true
	var boolean2 bool
	var error_ error
	var error_2 error = errors.New("Internal error")
	fmt.Println(boolean1, boolean2)
	fmt.Println(error_)
	fmt.Println(error_2)
}
