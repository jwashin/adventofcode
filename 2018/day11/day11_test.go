package main

import (
	"testing"
)

func Test_powerLevel(t *testing.T) {
	type args struct {
		x            int
		y            int
		serialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 5, 8}, 4},
		{"2", args{122, 79, 57}, -5},
		{"3", args{217, 196, 39}, 0},
		{"4", args{101, 153, 71}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerLevel(tt.args.x, tt.args.y, tt.args.serialNumber); got != tt.want {
				t.Errorf("powerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		serialNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{18}, "33,45"},
		{"2", args{42}, "21,61"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.serialNumber); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		serialNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{18}, "90,269,16"},
		// {"2", args{42}, "232,251,12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.serialNumber); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
