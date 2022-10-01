package main

import (
	"testing"
)

// {"1", args{`..#
// #..
// ...`, 10000}, 5587},

func Test_part1(t *testing.T) {
	type args struct {
		start         string
		activityCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2", args{`..#
#..
...`, 10000}, 5587},
		{"1", args{`..#
#..
...`, 70}, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diagnose(tt.args.start, tt.args.activityCount)
			if got != tt.want {
				t.Errorf("part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diagnose2(t *testing.T) {
	type args struct {
		start         string
		activityCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`..#
#..
...`, 100}, 26},
		{"2", args{`..#
#..
...`, 10000000}, 2511944},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagnose2(tt.args.start, tt.args.activityCount); got != tt.want {
				t.Errorf("diagnose2() = %v, want %v", got, tt.want)
			}
		})
	}
}
