package main

import (
	"testing"
)

func Test_gentestSpace(t *testing.T) {
	type args struct {
		factor int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{10}, `  0123456789
0 .#.####.##
1 ..#..#...#
2 #....##...
3 ###.#.###.
4 .##..#..#.
5 ..##....#.
6 #...##.###`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gentestSpace(tt.args.factor); got != tt.want {
				t.Errorf("gentestSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findShortestPath(t *testing.T) {
	type args struct {
		startx int
		starty int
		destx  int
		desty  int
		factor int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1, 7, 4, 10}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestPath(tt.args.startx, tt.args.starty, tt.args.destx, tt.args.desty, tt.args.factor); got != tt.want {
				t.Errorf("findShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
