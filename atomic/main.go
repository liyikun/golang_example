package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// has lock
	var wg1 sync.WaitGroup

	var ops1 uint64

	for i := 0; i < 50; i++ {
		wg1.Add(1)
		go func() {
			for i := 0; i < 5000; i++ {
				atomic.AddUint64(&ops1, 1)
			}
			wg1.Done()
		}()
	}

	wg1.Wait()

	// no lock
	var wg2 sync.WaitGroup

	var ops2 uint64

	for i := 0; i < 50; i++ {
		wg2.Add(1)
		go func() {
			for i := 0; i < 5000; i++ {
				ops2 += 1
			}
			wg2.Done()
		}()
	}

	wg2.Wait()

	fmt.Println("ops1 ", ops1, "ops2 ", ops2)
}
