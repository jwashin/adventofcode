package main

import (
	"testing"
)

func Test_swapPosition(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcde", 4, 0}, "ebcda"},
		{"2", args{"abcde", 3, 2}, "abdce"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPosition(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("swapPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapLetters(t *testing.T) {
	type args struct {
		s string
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"ebcda", "d", "b"}, "edcba"},
		{"2", args{"ebcda", "e", "a"}, "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapLetters(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("swapLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePositions(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"edcba", 0, 4}, "abcde"},
		{"2", args{"edcba", 1, 3}, "ebcda"},
		{"3", args{"edcba", 0, 3}, "bcdea"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePositions(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("reversePositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movePosition(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"bcdea", 1, 4}, "bdeac"},
		{"2", args{"bdeac", 3, 0}, "abdec"},
		{"3", args{"bdeac", 3, 4}, "bdeca"},
		{"4", args{"bdeac", 3, 2}, "bdaec"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movePosition(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("movePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateLeft(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcd", 1}, "bcda"},
		{"2", args{"abcd", 2}, "cdab"},
		{"3", args{"abcd", 0}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateLeft(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("rotateLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcd", 1}, "dabc"},
		{"2", args{"abcd", 2}, "cdab"},
		{"3", args{"abcd", 0}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateBasedOn(t *testing.T) {
	type args struct {
		s string
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abdec", "b"}, "ecabd"},
		{"2", args{"ecabd", "d"}, "decab"},
		// {"3", args{"abcdefgh", "d"}, "efghabcd"},
		// {"4", args{"abcdefgh", "a"}, "abcdefgh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateBasedOn(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("rotateBasedOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
