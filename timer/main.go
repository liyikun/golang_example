package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(2 * time.Second)

	fmt.Println("start")

	<-time1.C

	fmt.Println("end")

	time2 := time.NewTimer(2 * time.Second)

	go func() {
		<-time2.C
		fmt.Println("time2 end")
	}()

	stop2 := time2.Stop()

	if stop2 {
		fmt.Println("time2 stop end")
	}

	time.Sleep(3 * time.Second)
}
