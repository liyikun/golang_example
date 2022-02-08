package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map: ", m)

	v := m["k1"]

	fmt.Println("v: ", v)

	delete(m, "k2")

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("n: ", n)
}
