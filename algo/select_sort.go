package algo

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
