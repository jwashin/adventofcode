package main

import (
	"reflect"
	"testing"
)

// {"1", args{[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 20}, [][]int{{9, 11}}},

func Test_findCombos(t *testing.T) {
	type args struct {
		data  []int
		total int
		count int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 20, 2}, [][]int{{9, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCombos(tt.args.data, tt.args.total, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCombos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeLists(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`1
2
3
4
5
7
8
9
10
11`}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLists(tt.args.input); got != tt.want {
				t.Errorf("makeLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
