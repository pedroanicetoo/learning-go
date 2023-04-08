package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	public_place string
	number       uint8
}

func main() {
	var u user
	u.name = "Bar"
	u.age = 10
	fmt.Println(u)

	address_example := address{"street", 0}
	// also
	u2 := user{"Foo", 20, address_example}
	fmt.Println(u2)
	// also
	u3 := user{name: "Barz"}
	fmt.Println(u3)
}
