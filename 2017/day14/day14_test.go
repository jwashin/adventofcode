package main

import (
	"testing"
)

func Test_hash2binaryString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"a0c20170"}, "10100000110000100000000101110000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash2binaryString(tt.args.s); got != tt.want {
				t.Errorf("hash2binaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{true}, 8108},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.test); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		test bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{true}, 1242},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.test); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
