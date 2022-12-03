package main

import (
	"testing"
)

func Test_gamePlay(t *testing.T) {
	type args struct {
		opponent string
		player   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{"A", "Y"}, 8},
		{"2.", args{"B", "X"}, 1},
		{"3.", args{"C", "Z"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gamePlay(tt.args.opponent, tt.args.player); got != tt.want {
				t.Errorf("gamePlay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playGame(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`A Y
		B X
		C Z`}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.s); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamePlay2(t *testing.T) {
	type args struct {
		opponent string
		player   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{"A", "Y"}, 4},
		{"2.", args{"B", "X"}, 1},
		{"3.", args{"C", "Z"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gamePlay2(tt.args.opponent, tt.args.player); got != tt.want {
				t.Errorf("gamePlay2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playGame2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`A Y
		B X
		C Z`}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame2(tt.args.s); got != tt.want {
				t.Errorf("playGame2() = %v, want %v", got, tt.want)
			}
		})
	}
}
