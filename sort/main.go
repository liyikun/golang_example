package main

import (
	"fmt"
	"sort"
)

type byLength []string

// Len is the number of elements in the collection.
func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Less(i int, j int) bool {
	return len(b[i]) < len(b[j])
}

// Swap swaps the elements with indexes i and j.
func (b byLength) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	strs := []string{
		"c", "a", "d", "f",
	}

	sort.Strings(strs)

	fmt.Println("result:", strs)

	nums := []int{3, 1, 2, 6}

	sort.Ints(nums)

	fmt.Println("result:", nums)

	str2 := []string{
		"aae", "ee", "ccddd", "a",
	}

	sort.Sort(byLength(str2))

	fmt.Println(str2)

}
