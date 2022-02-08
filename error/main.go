package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		fmt.Println("read file error")
		log.Fatal(err)
	}
	fmt.Println(f)
	fmt.Println("hello")
}
