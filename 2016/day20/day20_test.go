package main

import (
	"testing"
)

func Test_lowestValAllowed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`5-8
0-2
4-7`}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestValAllowed(tt.args.s); got != tt.want {
				t.Errorf("lowestValAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countAllowed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`5-8
		0-2
		4-7`}, 4294967288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAllowed(tt.args.s); got != tt.want {
				t.Errorf("countAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
