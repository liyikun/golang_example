package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs

		fmt.Println("process exit", sig)

		done <- true

	}()

	go func() {
		<-time.After(2 * time.Second)
		os.Exit(3)
		done <- true
	}()

	<-done
}
