package main

import "fmt"

type person struct {
	name      string
	last_name string
	age       uint8
	height    float32
}

type student struct {
	person
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance:")

	p1 := person{"John", "Peter", 20, 1.72}
	fmt.Println(p1)
	p2 := student{p1, "History", "Oxford"}
	fmt.Println(p2)
	fmt.Println(p2.course)
	fmt.Println(p2.name)
}
