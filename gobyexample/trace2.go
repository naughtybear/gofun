package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()

	m := make(map[int]string, 10)
	m[1] = "lise"
	fmt.Println(m)

	n := map[string]int{"lisi": 10, "zhangsan": 2}
	fmt.Println(n)

	m = map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
	fmt.Println(m)
}
