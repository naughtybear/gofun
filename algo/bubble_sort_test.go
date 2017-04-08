package algo

import (
	"testing"
	"fmt"
)

func TestBubbleSort1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test bubble sort1",
			args{[]int{1, 2, 5, 3, 4, 2, 8, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort1(tt.args.arr)
		})
	}
}

func TestBubbleSort2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test bubble sort2",
			args{[]int{1, 2, 5, 3, 4, 2, 8, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort2(tt.args.arr)
			fmt.Println(tt.args)
		})
	}
}

func Test_bubbleSort3(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test bubble sort3",
			args{[]int{1, 2, 5, 3, 4, 2, 8, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort3(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
