package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string
	array1[0] = "index 0"
	fmt.Println(array1) // [index 0    ]

	array2 := [5]string{"index0", "index1", "index2", "index3", "index4"}
	fmt.Println(array2) // [index0 index1 index2 index3 index4]

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3) // [1 2 3 4 5]

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice) // [10 11 12 13 14 15 16 17]

	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array3)) // [5]int

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2) // [index1 index2]

	array2[1] = "Index1Change"
	fmt.Println(slice2) // [Index1Change index2]

	// Internal Arrays
	fmt.Println("---------------Internal Arrays---------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3) // [0 0 0 0 0 0 0 0 0 0]
	// length
	fmt.Println(len(slice3)) // 10
	// capacity
	fmt.Println(cap(slice3)) // 11

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)

	fmt.Println(slice3)      // [0 0 0 0 0 0 0 0 0 0 5 5]
	fmt.Println(len(slice3)) // 12
	fmt.Println(cap(slice3)) // 24

	slice4 := make([]float32, 5)
	fmt.Println(slice4) // [0 0 0 0 0]
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // 6
	fmt.Println(cap(slice4)) // 12
}
