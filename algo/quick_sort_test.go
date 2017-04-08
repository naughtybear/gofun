package algo

import "testing"
import "fmt"

func TestQuickSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	arr := []int{1, 2, 5, 3, 4, 2, 8, 6}
	tests := []struct {
		name string
		args args
	}{
		{
			"quick sort",
			args{arr, 0, len(arr) - 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr, tt.args.left, tt.args.right)
			fmt.Println(tt.args.arr)
		})
	}
}
