package algo

import "testing"

func Test_insertSort(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"insert1",
			args{[]int{1, 2, 5, 3, 4, 2, 8, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("before: ", tt.args.arr)
			insertSort(tt.args.arr)
			t.Log("after: ", tt.args.arr)
		})
	}
}

func Test_insertSort2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"insert2",
			args{[]int{1, 2, 5, 3, 4, 2, 8, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("before: ", tt.args.arr)
			insertSort2(tt.args.arr)
			t.Log("after: ", tt.args.arr)
		})
	}
}
