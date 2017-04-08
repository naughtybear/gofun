package algo

import "testing"
import "fmt"

func TestMergeSort(t *testing.T) {
	type args struct {
		arr   []int
		first int
		last  int
	}
	arr := []int{1, 2, 5, 3, 4, 2, 8, 6}
	tests := []struct {
		name string
		args args
	}{
		{"merge sort", args{arr, 0, len(arr) - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr, tt.args.first, tt.args.last)
			fmt.Println(tt.args.arr)
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		arr   []int
		first int
		mid   int
		last  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"merge", args{[]int{1, 2, 3, 2, 3, 4}, 0, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.args.arr, tt.args.first, tt.args.mid, tt.args.last)
			fmt.Println(tt.args.arr)
		})
	}
}
