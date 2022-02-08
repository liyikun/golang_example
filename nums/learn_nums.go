package main

import "fmt"

func example(num ...int) {

	for _, v := range num {
		fmt.Println(v)
	}

}

func main() {

	example(1, 2, 3, 4, 5)
}
