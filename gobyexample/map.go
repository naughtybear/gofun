package main

import (
	"fmt"
)

func main() {
	mymap := make(map[int]string, 2)
	mymap[1] = "a"
	mymap[2] = "b"
	mymap[3] = "c"

	fmt.Println("#" + mymap[4])

	if v, found := mymap[4]; found {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
}
