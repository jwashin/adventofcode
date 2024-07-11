package main

import (
	"testing"
)

func Test_firstAndLastDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{

		// TODO: Add test cases.
		{"1", args{"1abc2"}, 1, 2},
		{"2", args{"pqr3stu8vwx"}, 3, 8},
		{"3", args{"a1b2c3d4e5f"}, 1, 5},
		{"4", args{"treb7uchet"}, 7, 7},

		{"5", args{"two1nine"}, 2, 9},
		{"6", args{"eightwothree"}, 8, 3},
		{"7", args{"abcone2threexyz"}, 1, 3},
		{"8", args{"xtwone3four"}, 2, 4},
		{"9", args{"4nineeightseven2"}, 4, 2},
		{"10", args{"zoneight234"}, 1, 4},
		{"11", args{"7pqrstsixteen"}, 7, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := firstAndLastDigits(tt.args.s)
			if got != tt.want {
				t.Errorf("firstAndLastDigits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("firstAndLastDigits() got1 = %v, want %v", got1, tt.want1)
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

		{"1", args{`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`}, 142},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
