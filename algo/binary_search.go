package algo

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 22, 32}
// 	fmt.Println(binarySearch(arr, 7))
// 	fmt.Println(binarySearch(arr, 8))
// 	fmt.Println(binarySearch(arr, 4))
// 	fmt.Println(binarySearch(arr, 22))
// }

// BinarySearch binary search algo
func BinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)

	for i < j {
		k := (i + j) / 2
		if arr[k] == x {
			return k
		} else if arr[k] < x {
			i = k + 1
		} else {
			j = k - 1
		}
	}

	return -1
}
