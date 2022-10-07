package main

import (
	"testing"
)

// {"1", fields{circle: []int{0}, currentIndex: 0}, args{1}},
// 		{"2", fields{circle: []int{0, 1}, currentIndex: 1}, args{2}},

func Test_part1(t *testing.T) {
	type args struct {
		players int
		hiValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{9, 25}, 32},
		{"2", args{10, 1618}, 8317},
		{"3", args{13, 7999}, 146373},
		{"4", args{17, 1104}, 2764},
		{"5", args{21, 6111}, 54718},
		{"6", args{30, 5807}, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.players, tt.args.hiValue); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1a(t *testing.T) {
	type args struct {
		players int
		hiValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{9, 25}, 32},
		{"2", args{10, 1618}, 8317},
		{"3", args{13, 7999}, 146373},
		{"4", args{17, 1104}, 2764},
		{"5", args{21, 6111}, 54718},
		{"6", args{30, 5807}, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1a(tt.args.players, tt.args.hiValue); got != tt.want {
				t.Errorf("part1a() = %v, want %v", got, tt.want)
			}
		})
	}
}
