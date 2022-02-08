package main

import "fmt"

func test_chan_struct() {
	fmt.Println("start test chan struct ")
	type test_struct struct {
		name string
		age  int
	}
	i := test_struct{
		name: "li",
		age:  29,
	}

	// test chan

	ch1 := make(chan test_struct)

	go func() {
		ch1 <- i
	}()

	r1 := <-ch1

	fmt.Printf("value is %+v point %p \n", i, &i)

	fmt.Printf("value is %+v point %p \n", r1, &r1)

	r1.name = "l2"

	fmt.Printf("value is %+v point %p \n", r1, &r1)
	fmt.Printf("value is %+v point %p \n", i, &i)

	// struct 是完全 copy 值传递
}

func test_chan_map() {

	fmt.Println("start test chan map ")
	i := map[string]string{
		"name": "li",
		"test": "luo",
	}
	// test chan

	ch1 := make(chan map[string]string)

	go func() {
		ch1 <- i
	}()

	r1 := <-ch1

	fmt.Printf("value is %+v point %p \n", i, &i)

	fmt.Printf("value is %+v point %p \n", r1, &r1)

	r1["name"] = "l2"

	fmt.Printf("value is %+v point %p \n", r1, &r1)
	fmt.Printf("value is %+v point %p \n", i, &i)

	// map 现象是引用传递
}

func test_slice_chan() {

	fmt.Println("start test chan slice ")
	i := []int{
		1, 2, 3,
	}
	// test chan

	ch1 := make(chan []int)

	go func() {
		ch1 <- i
	}()

	r1 := <-ch1

	fmt.Printf("value is %+v point %p \n", i, &i)

	fmt.Printf("value is %+v point %p \n", r1, &r1)

	r1[1] = 10

	fmt.Printf("value is %+v point %p \n", r1, &r1)
	fmt.Printf("value is %+v point %p \n", i, &i)

	// slice 现象是引用传递
}

func test_chan_chan() {
	fmt.Println("start test chan chan ")
	i := make(chan int)

	go func() {
		i <- 10
	}()

	ch1 := make(chan chan int)

	go func() {
		ch1 <- i
	}()

	r1 := <-ch1

	fmt.Printf("value is %+v point %p \n", i, &i)

	fmt.Printf("value is %+v point %p \n", r1, &r1)

	v := <-r1

	fmt.Println("value is", v)

	fmt.Printf("value is %+v point %p \n", r1, &r1)
	fmt.Printf("value is %+v point %p \n", i, &i)

	// chan 现象是引用传递
}

func main() {
	test_chan_struct()

	test_chan_map()

	test_slice_chan()

	test_chan_chan()
}
