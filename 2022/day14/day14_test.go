package main

import (
	"testing"
)

func Test_dropSand(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`498,4 -> 498,6 -> 496,6
		503,4 -> 502,4 -> 502,9 -> 494,9`}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dropSand(tt.args.s); got != tt.want {
				t.Errorf("dropSand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dropSand2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`498,4 -> 498,6 -> 496,6
		503,4 -> 502,4 -> 502,9 -> 494,9`}, 93},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dropSand2(tt.args.s); got != tt.want {
				t.Errorf("dropSand2() = %v, want %v", got, tt.want)
			}
		})
	}
}
