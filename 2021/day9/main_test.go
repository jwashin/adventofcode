package main

import (
	"testing"
)

func Test_sumRisk(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`2199943210
		3987894921
		9856789892
		8767896789
		9899965678`}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRisk(tt.args.aString); got != tt.want {
				t.Errorf("sumRisk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doBasinTask(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`2199943210
		3987894921
		9856789892
		8767896789
		9899965678`}, 1134},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doBasinTask(tt.args.aString); got != tt.want {
				t.Errorf("doBasinTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
