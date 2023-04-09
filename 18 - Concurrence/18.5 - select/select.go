package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		case msgChannel1 := <-channel1:
			fmt.Println(msgChannel1)
		case msgChannel2 := <-channel2:
			fmt.Println(msgChannel2)
		}
	}

	/* using select we can provide to execute each channel without wasting time waiting, in this case channel1 executes each time.Millisecond * 500
	   while channel2 will be executed each time.Second * 2
	Channel 1
	Channel 1
	Channel 1
	Channel 2
	Channel 1
	Channel 1
	Channel 1
	Channel 1
	*/
}
