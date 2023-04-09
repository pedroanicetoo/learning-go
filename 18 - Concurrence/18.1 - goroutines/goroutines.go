package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World!") // goroutine
	write("Programing in Go")
	// Hello World!
	// Programing in Go
	// Programing in Go
	// Hello World!
	// ...
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
