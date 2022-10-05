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
		want string
	}{
		{"1", args{`Step C must be finished before step A can begin.
		Step C must be finished before step F can begin.
		Step A must be finished before step B can begin.
		Step A must be finished before step D can begin.
		Step B must be finished before step E can begin.
		Step D must be finished before step E can begin.
		Step F must be finished before step E can begin.`}, "CABDFE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderSteps(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSleigh(t *testing.T) {
	type args struct {
		s     string
		elves int
		test  bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Step C must be finished before step A can begin.
		Step C must be finished before step F can begin.
		Step A must be finished before step B can begin.
		Step A must be finished before step D can begin.
		Step B must be finished before step E can begin.
		Step D must be finished before step E can begin.
		Step F must be finished before step E can begin.`, 2, true}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSleigh(tt.args.s, tt.args.elves, tt.args.test); got != tt.want {
				t.Errorf("makeSleigh() = %v, want %v", got, tt.want)
			}
		})
	}
}
