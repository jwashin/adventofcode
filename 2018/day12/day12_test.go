package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		s           string
		generations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{`initial state: #..#.#..##......###...###

		...## => #
		..#.. => #
		.#... => #
		.#.#. => #
		.#.## => #
		.##.. => #
		.#### => #
		#.#.# => #
		#.### => #
		##.#. => #
		##.## => #
		###.. => #
		###.# => #
		####. => #`, 20}, 325},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s, tt.args.generations); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
