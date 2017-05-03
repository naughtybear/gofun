package main

import (
	"fmt"
	"math"
)

func main() {
	const factor = 3
	i := 20000
	i *= factor
	j := int16(20)
	i += int(j)
	fmt.Println(i)

	k := uint8(8)
	k = uint8(i) // 转换成功, 但是k的值为截为8位
	fmt.Println(k)

	fmt.Println(math.Gamma(4))
}

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x < math.MaxInt8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
