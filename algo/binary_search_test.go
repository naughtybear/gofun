package algo

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"search 7",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 22, 32}, 7},
			6,
		}, {
			"search 22",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 22, 32}, 22},
			8,
		}, {
			"search 4",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 22, 32}, 4},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
