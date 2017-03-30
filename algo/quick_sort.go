package algo

// func main() {
// 	arr := []int{1, 2, 5, 3, 4, 2, 8, 6}
// 	quickSort(arr, 0, len(arr)-1)
// 	fmt.Println(arr)
// }

func quickSort(arr []int, left, right int) {
	if left < right {
		k := arr[left]
		i, j := left, right

		for i < j {
			for i < j && arr[j] >= k {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for i < j && arr[i] < k {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = k
		quickSort(arr, left, i-1)
		quickSort(arr, i+1, right)
	}
}
