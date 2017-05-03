package main

import (
	"fmt"
)

func main() {
	for pos, char := range "日本\x80语" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)

	}

	a := "lisi"
	a := "zhangsan"
}
