package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String") // String
	generic(1)        // 1
	generic(true)     // true

	mapp := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapp)
}
