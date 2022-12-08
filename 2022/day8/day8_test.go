package main

import (
	"testing"
)

func Test_countVisible(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`30373
		25512
		65332
		33549
		35390`}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVisible(tt.args.s); got != tt.want {
				t.Errorf("countVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScenicScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`30373
		25512
		65332
		33549
		35390`}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScenicScore(tt.args.s); got != tt.want {
				t.Errorf("maxScenicScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
