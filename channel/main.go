package main

import "fmt"

func main() {
	message := make(chan string, 2)

	go func() {
		message <- "hello"
		message <- "test"
		message <- "test3"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
}
