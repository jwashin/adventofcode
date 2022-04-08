package main

import "testing"

func Test_doFirstInstruction(t *testing.T) {
	type args struct {
		dataFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{`6,10
		0,14
		9,10
		0,3
		10,4
		4,11
		6,0
		6,12
		4,1
		0,13
		10,12
		3,4
		3,0
		8,4
		1,10
		2,14
		8,10
		9,0
		
		fold along y=7
		fold along x=5`}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doFirstInstruction(tt.args.dataFile); got != tt.want {
				t.Errorf("doFirstInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
