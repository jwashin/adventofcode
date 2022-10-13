package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{9}, "5158916779"},
		{"2", args{5}, "0124515891"},
		{"5", args{2018}, "5941429882"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.index); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		index string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5", args{"59414"}, 2018},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.index); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
