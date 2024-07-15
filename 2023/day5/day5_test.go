package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`}, 35},
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
		{"1", args{`seeds: 79 14 55 13

		seed-to-soil map:
		50 98 2
		52 50 48
		
		soil-to-fertilizer map:
		0 15 37
		37 52 2
		39 0 15
		
		fertilizer-to-water map:
		49 53 8
		0 11 42
		42 0 7
		57 7 4
		
		water-to-light map:
		88 18 7
		18 25 70
		
		light-to-temperature map:
		45 77 23
		81 45 19
		68 64 13
		
		temperature-to-humidity map:
		0 69 1
		1 0 69
		
		humidity-to-location map:
		60 56 37
		56 93 4`}, 46}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
