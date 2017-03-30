package algo

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 5, 3, 4, 2, 8, 6}
	bubbleSort1(arr)
	fmt.Println(arr)

	arr = []int{1, 2, 5, 3, 4, 2, 8, 6}
	bubbleSort2(arr)
	fmt.Println(arr)

	arr = []int{1, 2, 5, 3, 4, 2, 8, 6}
	bubbleSort3(arr)
	fmt.Println(arr)
}

func bubbleSort1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func bubbleSort2(arr []int) {
	swap := true
	i := 0
	for swap {
		swap = false
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swap = true
			}
		}
		i++
	}
}

func bubbleSort3(arr []int) {
	i, k := 0, len(arr)
	flag := k
	for flag > 0 {
		k = flag
		flag = 0
		for j := i + 1; j < k; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = j
			}
		}
	}
}
