package main

import (
	"fmt"
	"sync"
)

type data struct {
	v map[string]int
	m sync.Mutex
}

func (d *data) inc(k string, id int) {
	d.m.Lock()
	defer d.m.Unlock()
	fmt.Println("lock ", id)
	d.v[k]++
}

func (d *data) value(k string) int {
	d.m.Lock()
	defer d.m.Unlock()
	fmt.Println("get k ", k)
	return d.v[k]
}

func main() {
	var wg sync.WaitGroup

	d := data{
		v: make(map[string]int),
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		t := i
		go func() {
			d.inc("test", t)
			fmt.Println("value is", d.value("test"))
			defer wg.Done()
		}()
	}

	wg.Wait()

}
