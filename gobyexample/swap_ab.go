package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 5

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)

	a = 4
	b = 4

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
