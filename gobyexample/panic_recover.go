package main

import (
	"fmt"
)

func f() {
	fmt.Println("a exception happened...")
	panic("an error just happened...")
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("%v", x)
		}
	}()
	f()
}
