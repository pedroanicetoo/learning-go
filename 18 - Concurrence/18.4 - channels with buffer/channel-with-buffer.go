package main

import "fmt"

func main() {
	channel := make(chan string, 2) // will be block channel only if reach its maximum capacity which is 2

	channel <- "Hello World!"
	channel <- "Programing in Go"
	channel <- "This line will raises deadlock on channel cause size of buffer is 2"

	msg := <-channel
	msg2 := <-channel

	fmt.Println(msg)
	fmt.Println(msg2)
}
