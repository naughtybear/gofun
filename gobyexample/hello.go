package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	a := make(map[string]interface{})
	a["lisi"] = 3
	fmt.Println(a)

	b := map[string]interface{}{}
	b["zhangsan"] = "zhangsan"
	fmt.Println(b)
}
