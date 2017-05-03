package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (ss ByLength) Len() int {
	return len(ss)
}

func (ss ByLength) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss ByLength) Less(i, j int) bool {
	return len(ss[i]) < len(ss[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{2, 1, 8, 0}
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println("Sorted:", sort.IntsAreSorted(ints))

	// sort by func
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
