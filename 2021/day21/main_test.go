package main

import (
	"testing"
)

// {"1", args{1, 1}, 2},
// 		{"2", args{10, 1}, 1},
// 		{"3", args{7, 5}, 2},
// 		{"4", args{6, 4}, 10},

// {"1", args{`Player 1 starting position: 4
// 		Player 2 starting position: 8`}, 739785},

// {"1", args{`Player 1 starting position: 4
// Player 2 starting position: 8`, 21}, 444356092776315},

func Test_start(t *testing.T) {
	type args struct {
		aString      string
		winningScore uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{`Player 1 starting position: 4
		Player 2 starting position: 8`, 1000}, 739785},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := start(tt.args.aString, tt.args.winningScore); got != tt.want {
				t.Errorf("start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock10Add(t *testing.T) {
	type args struct {
		start uint64
		anInt uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{1, 1}, 2},
		{"2", args{10, 1}, 1},
		{"3", args{7, 5}, 2},
		{"4", args{6, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clock10Add(tt.args.start, tt.args.anInt); got != tt.want {
				t.Errorf("clock10Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dirac(t *testing.T) {
	type args struct {
		aString      string
		winningScore uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{`Player 1 starting position: 4
Player 2 starting position: 8`, 21}, 444356092776315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirac(tt.args.aString, tt.args.winningScore); got != tt.want {
				t.Errorf("dirac() = %v, want %v", got, tt.want)
			}
		})
	}
}
