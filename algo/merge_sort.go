package algo

// MergeSort this is a implemention
func MergeSort(arr []int, first, last int) {
	if first < last {
		mid := (first + last) / 2
		MergeSort(arr, first, mid)
		MergeSort(arr, mid+1, last)
		Merge(arr, first, mid, last)
	}
}

// Merge 合并两列已排序的数组
func Merge(arr []int, first, mid, last int) {
	tmp := make([]int, last-first+1)
	i, j := first, mid+1
	k := 0
	for i <= mid && j <= last {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= last {
		tmp[k] = arr[j]
		k++
		j++
	}
	for i := 0; i < k; i++ {
		arr[first+i] = tmp[i]
	}
}
