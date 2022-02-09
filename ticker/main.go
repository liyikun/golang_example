package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1000 * time.Microsecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("ticker at", t)
			}
		}
	}()

	time.Sleep(3200 * time.Microsecond)
	ticker.Stop()

	done <- true
	fmt.Println("finish")

	i := 0

	for v := range time.Tick(time.Second) {
		if i < 10 {
			fmt.Println(i)
			fmt.Println(v)
			i++
		} else {
			fmt.Println("finish")
			return
		}
	}
}
