package main

import (
	"testing"
)

func Test_surfaceArea(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"0", args{`1,1,1
		2,1,1`}, 10},

		{"1", args{`2,2,2
		1,2,2
		3,2,2
		2,1,2
		2,3,2
		2,2,1
		2,2,3
		2,2,4
		2,2,6
		1,2,5
		3,2,5
		2,1,5
		2,3,5`}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := surfaceArea(tt.args.s); got != tt.want {
				t.Errorf("surfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_externalSurfaceArea(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`2,2,2
		1,2,2
		3,2,2
		2,1,2
		2,3,2
		2,2,1
		2,2,3
		2,2,4
		2,2,6
		1,2,5
		3,2,5
		2,1,5
		2,3,5`}, 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := externalSurfaceArea(tt.args.s); got != tt.want {
				t.Errorf("externalSurfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
