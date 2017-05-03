package main

import (
	"fmt"
)

func main() {
	a, b, c := 1, 2, 3
	for a := 7; a < 8; a++ { // a无意间覆盖了外部的a
		b := 7 // b无意间覆盖了外部的b
		c = 13 // c为外部的c值
		fmt.Println("inner:", a, b, c)
	}
	fmt.Println("outer:", a, b, c)

	// just for test
	fancy := StringSlice([]string{"a", "b"})
	fmt.Println(fancy)
}

type StringSlice []string
