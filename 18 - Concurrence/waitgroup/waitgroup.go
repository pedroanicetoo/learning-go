package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello World!")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Programing in Go")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0

	/* OUTPUT:
	Programing in Go
	Hello World!
	Hello World!
	Programing in Go
	Hello World!
	Programing in Go
	Programing in Go
	Hello World!
	Hello World!
	Programing in Go
	*/

}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
