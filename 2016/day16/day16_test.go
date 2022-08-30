package main

import (
	"testing"
)

func Test_filldisk(t *testing.T) {
	type args struct {
		init   string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"10000", 20}, "01100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filldisk(tt.args.init, tt.args.length); got != tt.want {
				t.Errorf("filldisk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dragonCurve2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1"}, "100"},
		{"2", args{"0"}, "001"},
		{"3", args{"11111"}, "11111000000"},
		{"4", args{"111100001010"}, "1111000010100101011110000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dragonCurve2(tt.args.s); got != tt.want {
				t.Errorf("dragonCurve2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSum2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"110010110100"}, "100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSum2(tt.args.s); got != tt.want {
				t.Errorf("checkSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
