package main

import (
	"reflect"
	"testing"
)

// 		{"1", args{`###########
// #0.1.....2#
// #.#######.#
// #4.......3#
// ###########
// `}, map[int]string{0: "1,1", 1: "3,1", 2: "9,1", 3: "9,3", 4: "1,3"}},

func Test_charGrid_getPointsOfInterest(t *testing.T) {
	tests := []struct {
		name string
		c    charGrid
		want map[int]string
	}{
		{"1", charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}, map[int]string{0: "1,1", 1: "3,1", 2: "9,1", 3: "9,3", 4: "1,3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getPointsOfInterest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charGrid.getPointsOfInterest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charGrid_getShortestPath(t *testing.T) {
	type args struct {
		start string
		dest  string
	}
	tests := []struct {
		name string
		c    charGrid
		args args
		want int
	}{
		{"0 to 4", charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}, args{"1,1", "1,3"}, 2},
		{"4 to 1", charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}, args{"1,3", "3,1"}, 4},
		{"1 to 2", charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}, args{"3,1", "9,1"}, 6},
		{"2 to 3", charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}, args{"9,1", "9,3"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getShortestPath(tt.args.start, tt.args.dest); got != tt.want {
				t.Errorf("charGrid.getShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		grid charGrid
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{charGrid{"###########", "#0.1.....2#", "#.#######.#", "#4.......3#", "###########"}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.grid); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
