package main

import (
	"reflect"
	"testing"
)

func Test_test(t *testing.T) {
	type args struct {
		incrementCount int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"1", args{0}, 0, []int{0}},
		{"2", args{1}, 1, []int{0, 1}},
		{"3", args{2}, 1, []int{0, 2, 1}},
		{"4", args{3}, 2, []int{0, 2, 3, 1}},
		{"9", args{9}, 1, []int{0, 9, 5, 7, 2, 4, 3, 8, 6, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := test(tt.args.incrementCount)
			if got != tt.want {
				t.Errorf("test() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("test() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{true}, 638},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.test); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
