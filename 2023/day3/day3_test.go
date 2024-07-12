package main

import (
	"reflect"
	"testing"
)

func Test_getValidNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1.", args{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`}, []int{467, 35, 633, 617, 592, 755, 664, 598}},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValidNumbers(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValidNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
