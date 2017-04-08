package algo

import "testing"

func Test_selectSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"select sort", args{[]int{1, 2, 5, 3, 4, 2, 8, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectSort(tt.args.arr)
			want := []int{1, 2, 2, 3, 4, 5, 6, 8}
			for i, k := range want {
				if tt.args.arr[i] != k {
					t.Errorf("selectSort() = %v, want %v", tt.args.arr, want)
				}
			}
		})
	}
}
