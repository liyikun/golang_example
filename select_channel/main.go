package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}
	}
}
