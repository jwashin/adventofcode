package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`#######   
		#.G...#
		#...EG#
		#.#.#G#
		#..G#E#
		#.....#   
		#######`}, 27730},
		{"2", args{`#######
		#G..#E#
		#E#E.E#
		#G.##.#
		#...#E#
		#...E.#
		#######`}, 36334},
		{"3", args{`#######
		#E..EG#
		#.#G.E#
		#E.##E#
		#G..#.#
		#..E#.#
		#######  `}, 39514},
		{"4", args{`#######
		#E.G#.#
		#.#G..#
		#G.#.G#
		#G..#.#
		#...E.#
		#######  `}, 27755},
		{"5", args{`#######
		#.E...#
		#.#..G#
		#.###.#
		#E#G#G#
		#...#G#
		####### `}, 28944},
		{"6", args{`#########
		#G......#
		#.E.#...#
		#..##..G#
		#...##..#
		#...#...#
		#.G...G.#
		#.....G.#
		#########  `}, 18740},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
