package main

import (
	"fmt"
)

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("hell a")
}

func main() {
	a()
}
