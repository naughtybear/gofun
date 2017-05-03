package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println("total =", total)
}

func main() {
	sum(1, 2, 3, 4, 5)

	sum([]int{1, 2, 3, 4, 5}...)
}
