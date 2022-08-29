package main

import (
	"testing"
)

func Test_clockmath(t *testing.T) {
	type args struct {
		nPositions int
		position   int
		addend     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 4, 1}, 0},
		{"2", args{2, 1, 2}, 1},
		{"3", args{5, 4, 6}, 0},
		{"4", args{2, 1, 7}, 0},
		{"5", args{5, 0, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clockmath(tt.args.nPositions, tt.args.position, tt.args.addend); got != tt.want {
				t.Errorf("clockmath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{"1", args{"Disc #1 has 5 positions; at time=0, it is at position 4."}, 1, 5, 4},
		{"2", args{"Disc #2 has 2 positions; at time=0, it is at position 1."}, 2, 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := parseInput(tt.args.s)
			if got != tt.want {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseInput() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_doPush(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPush(tt.args.s); got != tt.want {
				t.Errorf("doPush() = %v, want %v", got, tt.want)
			}
		})
	}
}
