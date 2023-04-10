package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexing(write("Hello World !"), write("Programing in GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexing(entryChannel1, entryChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-entryChannel1:
				outputChannel <- message
			case message := <-entryChannel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received Value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
		}
	}()

	return channel
}
