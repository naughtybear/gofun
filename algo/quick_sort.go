package algo

// QuickSort quick sort
func QuickSort(arr []int, left, right int) {
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
		QuickSort(arr, left, i-1)
		QuickSort(arr, i+1, right)
	}
}
