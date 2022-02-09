package main

import "fmt"

func myPanic() {
	panic("message")
}

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("error is ", e)
		}
	}()

	myPanic()

	fmt.Println("finish")
}
