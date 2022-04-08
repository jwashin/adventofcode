package main

import (
	"testing"
)

func TestLeastFuelSpent(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"16,1,2,0,4,2,7,1,2,14"}, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeastFuelSpent(tt.args.data); got != tt.want {
				t.Errorf("LeastFuelSpent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_triangularCost(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 3},
		{"5", args{5}, 15},
		{"3", args{3}, 6},
		{"11", args{11}, 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularSum(tt.args.v); got != tt.want {
				t.Errorf("triangularCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeastFuelSpent2(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"16,1,2,0,4,2,7,1,2,14"}, 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeastFuelSpent2(tt.args.data); got != tt.want {
				t.Errorf("LeastFuelSpent2() = %v, want %v", got, tt.want)
			}
		})
	}
}
