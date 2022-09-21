package main

import (
	"reflect"
	"testing"
)

func Test_cycle(t *testing.T) {
	type args struct {
		bank []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{0, 2, 7, 0}}, []int{2, 4, 1, 2}},
		{"2", args{[]int{2, 4, 1, 2}}, []int{3, 1, 2, 3}},
		{"3", args{[]int{3, 1, 2, 3}}, []int{0, 2, 3, 4}},
		{"4", args{[]int{0, 2, 3, 4}}, []int{1, 3, 4, 1}},
		{"5", args{[]int{1, 3, 4, 1}}, []int{2, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cycle(tt.args.bank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"0 2 7 0"}, 5},
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
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"0 2 7 0"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
