package main

import (
	"testing"
)

func Test_connectsToZero(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0 <-> 2
		1 <-> 1
		2 <-> 0, 3, 4
		3 <-> 2, 4
		4 <-> 2, 3, 6
		5 <-> 6
		6 <-> 4, 5`}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectsToZero(tt.args.s); got != tt.want {
				t.Errorf("connectsToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countGroups(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0 <-> 2
		1 <-> 1
		2 <-> 0, 3, 4
		3 <-> 2, 4
		4 <-> 2, 3, 6
		5 <-> 6
		6 <-> 4, 5`}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGroups(tt.args.s); got != tt.want {
				t.Errorf("countGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
