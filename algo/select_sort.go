package algo

// func main() {
// 	arr := []int{1, 2, 5, 3, 4, 2, 8, 6}
// 	selectSort(arr)
// 	fmt.Println(arr)
// }

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		nMinIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[nMinIdx] {
				nMinIdx = j
			}
		}
		if i != nMinIdx {
			arr[i], arr[nMinIdx] = arr[nMinIdx], arr[i]
		}
	}
}
