package algo

// BubbleSort1 bubble sort impletation 1
func BubbleSort1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// BubbleSort2 bubble sort impletation 2
func BubbleSort2(arr []int) {
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

// BubbleSort3 bubble sort impletation 3
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
