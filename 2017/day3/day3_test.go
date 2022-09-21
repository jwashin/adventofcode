package main

import (
	"testing"
)

func Test_leastSquareGE(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{3}, 3},
		{"3", args{11}, 5},
		{"4", args{10}, 5},
		{"6", args{9}, 3},
		{"5", args{28}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastSquareGE(tt.args.n); got != tt.want {
				t.Errorf("leastSquareGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coordfromX(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{9}, 1, 1},
		{"3", args{25}, 2, 2},
		{"4", args{3}, -1, 1},
		{"2", args{2}, 0, 1},
		{"5", args{20}, 1, -2},
		{"6", args{1}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := coordfromX(tt.args.n)
			if got != tt.want {
				t.Errorf("coordfromX() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("coordfromX() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_spiralDistance(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 0},
		{"12", args{12}, 3},
		{"3", args{23}, 2},
		{"4", args{1024}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralDistance(tt.args.s); got != tt.want {
				t.Errorf("spiralDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
