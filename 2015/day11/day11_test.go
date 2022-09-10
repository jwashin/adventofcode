package main

import (
	"testing"
)

func Test_increment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"xx"}, "xy"},
		{"2", args{"xy"}, "xz"},
		{"3", args{"xz"}, "ya"},
		{"4", args{"ya"}, "yb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increment(tt.args.s); got != tt.want {
				t.Errorf("increment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"hijklmmn"}, false},
		{"2", args{"abbceffg"}, false},
		{"3", args{"abbcegjk"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPassword(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcdefgh"}, "abcdffaa"},
		{"2", args{"ghijklmn"}, "ghjaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPassword(tt.args.s); got != tt.want {
				t.Errorf("nextPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
