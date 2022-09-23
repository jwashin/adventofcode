package main

import (
	"reflect"
	"testing"
)

func Test_knotHash(t *testing.T) {
	type args struct {
		lengths []int
		test    bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3}, true}, []int{2, 1, 0, 3, 4}},
		{"2", args{[]int{3, 4}, true}, []int{4, 3, 0, 1, 2}},
		{"3", args{[]int{3, 4, 1}, true}, []int{4, 3, 0, 1, 2}},
		{"4", args{[]int{3, 4, 1, 5}, true}, []int{3, 4, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knotHash(tt.args.lengths, tt.args.test); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("knotHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "a2582a3a0e66e6e86e3812dcb672a272"},
		{"2", args{"AoC 2017"}, "33efeb34ea91902bb2f59c9920caa6cd"},
		{"3", args{"1,2,3"}, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"4", args{"1,2,4"}, "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.data); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xorList(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorList(tt.args.s); got != tt.want {
				t.Errorf("xorList() = %v, want %v", got, tt.want)
			}
		})
	}
}
