package main

import "fmt"

func test_map() {

	test_info := map[string]string{
		"name": "li",
		"age":  "23",
	}

	fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	func(v map[string]string) {

		fmt.Printf("arguments v is %+v , point is %p \n", v, &v)

		v["name"] = "x"

		fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	}(test_info)

	// map 现象是引用传递
}

func test_struct() {

	type test_type struct {
		name string
		age  int
	}

	test_info := test_type{
		name: "liyikun",
		age:  29,
	}

	fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	func(v test_type) {

		fmt.Printf("arguments v is %+v , point is %p \n", v, &v)

		v.name = "x"

		fmt.Printf("arguments v is %+v , point is %p \n", v, &v)
		fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	}(test_info)

	// struct 是值传递，完全copy
}

func test_slice() {

	test_info := []int{1, 2, 3}

	fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	func(v []int) {

		fmt.Printf("arguments v is %v , point is %p \n", v, &v)

		v[1] = 10

		fmt.Println("default v is ", v)

	}(test_info)

	// slice 现象是引用传递
}

func test_chan() {

	test_info := make(chan int)

	fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	go func(v chan int) {

		fmt.Printf("arguments v is %+v , point is %p \n", v, &v)

		v <- 10

		fmt.Printf("default v is %+v , point is %p \n", test_info, &test_info)

	}(test_info)

	fmt.Println(<-test_info)
	// chan 现象是引用传递
}

func main() {
	test_map()
	test_struct()
	test_chan()
	test_slice()
}
