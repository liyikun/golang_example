package main

import (
	"fmt"
	"sync"
	"time"
)

func run(index int) {
	fmt.Printf("sleep %d start \n", index)
	time.Sleep(time.Second)
	fmt.Printf("sleep %d end \n", index)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			run(i)
		}()
	}

	wg.Wait()
}
