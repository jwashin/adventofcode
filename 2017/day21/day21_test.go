package main

import (
	"testing"
)

func Test_rotate90(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"../.#"}, "../#."},
		{"2", args{"../#."}, "#./.."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate90(tt.args.s); got != tt.want {
				t.Errorf("rotate90() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
