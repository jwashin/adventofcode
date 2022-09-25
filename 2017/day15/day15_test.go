package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{true}, 588},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.test); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{true}, 309},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.test); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
