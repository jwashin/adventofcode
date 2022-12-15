package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s   string
		row int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		Sensor at x=9, y=16: closest beacon is at x=10, y=16
		Sensor at x=13, y=2: closest beacon is at x=15, y=3
		Sensor at x=12, y=14: closest beacon is at x=10, y=16
		Sensor at x=10, y=20: closest beacon is at x=10, y=16
		Sensor at x=14, y=17: closest beacon is at x=10, y=16
		Sensor at x=8, y=7: closest beacon is at x=2, y=10
		Sensor at x=2, y=0: closest beacon is at x=2, y=10
		Sensor at x=0, y=11: closest beacon is at x=2, y=10
		Sensor at x=20, y=14: closest beacon is at x=25, y=17
		Sensor at x=17, y=20: closest beacon is at x=21, y=22
		Sensor at x=16, y=7: closest beacon is at x=15, y=3
		Sensor at x=14, y=3: closest beacon is at x=15, y=3
		Sensor at x=20, y=1: closest beacon is at x=15, y=3`, 10}, 26},
		{"2", args{`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		Sensor at x=9, y=16: closest beacon is at x=10, y=16
		Sensor at x=13, y=2: closest beacon is at x=15, y=3
		Sensor at x=12, y=14: closest beacon is at x=10, y=16
		Sensor at x=10, y=20: closest beacon is at x=10, y=16
		Sensor at x=14, y=17: closest beacon is at x=10, y=16
		Sensor at x=8, y=7: closest beacon is at x=2, y=10
		Sensor at x=2, y=0: closest beacon is at x=2, y=10
		Sensor at x=0, y=11: closest beacon is at x=2, y=10
		Sensor at x=20, y=14: closest beacon is at x=25, y=17
		Sensor at x=17, y=20: closest beacon is at x=21, y=22
		Sensor at x=16, y=7: closest beacon is at x=15, y=3
		Sensor at x=14, y=3: closest beacon is at x=15, y=3
		Sensor at x=20, y=1: closest beacon is at x=15, y=3`, 11}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s, tt.args.row); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		s   string
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		Sensor at x=9, y=16: closest beacon is at x=10, y=16
		Sensor at x=13, y=2: closest beacon is at x=15, y=3
		Sensor at x=12, y=14: closest beacon is at x=10, y=16
		Sensor at x=10, y=20: closest beacon is at x=10, y=16
		Sensor at x=14, y=17: closest beacon is at x=10, y=16
		Sensor at x=8, y=7: closest beacon is at x=2, y=10
		Sensor at x=2, y=0: closest beacon is at x=2, y=10
		Sensor at x=0, y=11: closest beacon is at x=2, y=10
		Sensor at x=20, y=14: closest beacon is at x=25, y=17
		Sensor at x=17, y=20: closest beacon is at x=21, y=22
		Sensor at x=16, y=7: closest beacon is at x=15, y=3
		Sensor at x=14, y=3: closest beacon is at x=15, y=3
		Sensor at x=20, y=1: closest beacon is at x=15, y=3`, 20}, 56000011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s, tt.args.max); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
// ###S# ##### ##### ##.## ##### ####