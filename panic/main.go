package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("defer")
	}()

	_, err := os.Open("wefew.txt")

	if err != nil {
		panic(err)
	}
}
