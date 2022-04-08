package main

import (
	"testing"
)

func Test_flashOctopi(t *testing.T) {
	type args struct {
		input string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{`11111
		19991
		19191
		19991
		11111`, 1}, 9},
		{"example2", args{`11111
		19991
		19191
		19991
		11111`, 2}, 9},
		{"example3", args{`5483143223
		2745854711
		5264556173
		6141336146
		6357385478
		4167524645
		2176841721
		6882881134
		4846848554
		5283751526`, 10}, 204},
		{"example3", args{`5483143223
		2745854711
		5264556173
		6141336146
		6357385478
		4167524645
		2176841721
		6882881134
		4846848554
		5283751526`, 100}, 1656},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flashOctopi(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("flashOctopi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstStepToUnisonFlash(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example3", args{`5483143223
		2745854711
		5264556173
		6141336146
		6357385478
		4167524645
		2176841721
		6882881134
		4846848554
		5283751526`}, 195},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStepToUnisonFlash(tt.args.input); got != tt.want {
				t.Errorf("firstStepToUnisonFlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
