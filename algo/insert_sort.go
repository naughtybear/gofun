package algo

func insertSort(arr []int) {
	i := 0
	for j := i + 1; j < len(arr); j++ {
		if arr[j] < arr[i] {
			for k := i; k > 0 && arr[k] > arr[j]; k-- { // arr[k] > arr[j] 这种方式是有问题的，arr[j]会因为数据交换变动
				arr[k+1], arr[k] = arr[k], arr[k+1]
			}
		}
		i++
	}
}

func insertSort2(arr []int) {
	i := 0
	for j := i + 1; j < len(arr); j++ {
		if arr[j] < arr[i] {
			x := arr[j] // x是必须的，否则在交换后arr[j]的值会发生改变，insertSort方法会发生这种情况，导致排序错误
			// i - 最有一个有序元素的位置
			for k := i; k > 0 && arr[k] > x; k-- {
				arr[k+1], arr[k] = arr[k], arr[k+1]
			}
		}
		i++
	}
}
