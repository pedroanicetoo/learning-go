package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Hello World !", channel)

	fmt.Println("After of write function starts executed...")

	// msg := <-channel // wait to channel receive value...
	// fmt.Println(msg)

	// for {
	// 	msg := <-channel
	// 	fmt.Println(msg)
	// } // after loop: fatal error: all goroutines are asleep - deadlock!
	// channel will be waiting for value but will be never happen, cause this deadlock is generated
	// To resolve this add close function to channel...

	// for {
	// 	msg, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	// ALSO
	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Program ending...")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
