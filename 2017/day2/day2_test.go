package main

import "testing"

func Test_xdivision(t *testing.T) {
	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 9, 2, 8}}, 4},
		{"2", args{[]int{9, 4, 7, 3}}, 3},
		{"3", args{[]int{3, 8, 6, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xdivision(tt.args.r); got != tt.want {
				t.Errorf("xdivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
