package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`1, 1
		1, 6
		8, 3
		3, 4
		5, 5
		8, 9`}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		s           string
		maxDistance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`1, 1
		1, 6
		8, 3
		3, 4
		5, 5
		8, 9`, 32}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s, tt.args.maxDistance); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
