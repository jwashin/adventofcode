package main

import "testing"

// {"1", args{`.#.#.#
// ...##.
// #....#
// ..#...
// #.#..#
// ####..`, 1}, `..##..
// ..##.#
// ...##.
// ......
// #.....
// #.##..`},

func Test_life(t *testing.T) {
	type args struct {
		s    string
		iter int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, 0}, `.#.#.#
...##.
#....#
..#...
#.#..#
####..`},

		{"1", args{`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, 1}, `..##..
..##.#
...##.
......
#.....
#.##..`},

		{"4", args{`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, 4}, `......
......
..##..
..##..
......
......`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := life(tt.args.s, tt.args.iter); got != tt.want {
				t.Errorf("life() = %v, want %v", got, tt.want)
			}
		})
	}
}
