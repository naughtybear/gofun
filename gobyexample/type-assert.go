package main

import (
	"fmt"
)

func main() {
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	j := i.(int) // j是int类型的数据(或者发生一个panic)
	fmt.Printf("%T->%d\n", j, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T->%d\n", i, i)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T->%q\n", s, s)
	}

	// 类型断言的类型必须是接口类型
}
