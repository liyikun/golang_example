package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.bytedance.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for {
		v := scanner.Scan()
		if v {
			fmt.Println(scanner.Text())
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
