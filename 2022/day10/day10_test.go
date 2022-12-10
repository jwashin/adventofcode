package main

import (
	"testing"
)

func Test_signalStrength(t *testing.T) {
	type args struct {
		s string
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`addx 15
		addx -11
		addx 6
		addx -3
		addx 5
		addx -1
		addx -8
		addx 13
		addx 4
		noop
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx -35
		addx 1
		addx 24
		addx -19
		addx 1
		addx 16
		addx -11
		noop
		noop
		addx 21
		addx -15
		noop
		noop
		addx -3
		addx 9
		addx 1
		addx -3
		addx 8
		addx 1
		addx 5
		noop
		noop
		noop
		noop
		noop
		addx -36
		noop
		addx 1
		addx 7
		noop
		noop
		noop
		addx 2
		addx 6
		noop
		noop
		noop
		noop
		noop
		addx 1
		noop
		noop
		addx 7
		addx 1
		noop
		addx -13
		addx 13
		addx 7
		noop
		addx 1
		addx -33
		noop
		noop
		noop
		addx 2
		noop
		noop
		noop
		addx 8
		noop
		addx -1
		addx 2
		addx 1
		noop
		addx 17
		addx -9
		addx 1
		addx 1
		addx -3
		addx 11
		noop
		noop
		addx 1
		noop
		addx 1
		noop
		noop
		addx -13
		addx -19
		addx 1
		addx 3
		addx 26
		addx -30
		addx 12
		addx -1
		addx 3
		addx 1
		noop
		noop
		noop
		addx -9
		addx 18
		addx 1
		addx 2
		noop
		noop
		addx 9
		noop
		noop
		noop
		addx -1
		addx 2
		addx -37
		addx 1
		addx 3
		noop
		addx 15
		addx -21
		addx 22
		addx -6
		addx 1
		noop
		addx 2
		addx 1
		noop
		addx -10
		noop
		noop
		addx 20
		addx 1
		addx 2
		addx 2
		addx -6
		addx -11
		noop
		noop
		noop`, 220}, 3960},
		{"1", args{`addx 15
		addx -11
		addx 6
		addx -3
		addx 5
		addx -1
		addx -8
		addx 13
		addx 4
		noop
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx -35
		addx 1
		addx 24
		addx -19
		addx 1
		addx 16
		addx -11
		noop
		noop
		addx 21
		addx -15
		noop
		noop
		addx -3
		addx 9
		addx 1
		addx -3
		addx 8
		addx 1
		addx 5
		noop
		noop
		noop
		noop
		noop
		addx -36
		noop
		addx 1
		addx 7
		noop
		noop
		noop
		addx 2
		addx 6
		noop
		noop
		noop
		noop
		noop
		addx 1
		noop
		noop
		addx 7
		addx 1
		noop
		addx -13
		addx 13
		addx 7
		noop
		addx 1
		addx -33
		noop
		noop
		noop
		addx 2
		noop
		noop
		noop
		addx 8
		noop
		addx -1
		addx 2
		addx 1
		noop
		addx 17
		addx -9
		addx 1
		addx 1
		addx -3
		addx 11
		noop
		noop
		addx 1
		noop
		addx 1
		noop
		noop
		addx -13
		addx -19
		addx 1
		addx 3
		addx 26
		addx -30
		addx 12
		addx -1
		addx 3
		addx 1
		noop
		noop
		noop
		addx -9
		addx 18
		addx 1
		addx 2
		noop
		noop
		addx 9
		noop
		noop
		noop
		addx -1
		addx 2
		addx -37
		addx 1
		addx 3
		noop
		addx 15
		addx -21
		addx 22
		addx -6
		addx 1
		noop
		addx 2
		addx 1
		noop
		addx -10
		noop
		noop
		addx 20
		addx 1
		addx 2
		addx 2
		addx -6
		addx -11
		noop
		noop
		noop`, 60}, 1140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signalStrength(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("signalStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`addx 15
		addx -11
		addx 6
		addx -3
		addx 5
		addx -1
		addx -8
		addx 13
		addx 4
		noop
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx -35
		addx 1
		addx 24
		addx -19
		addx 1
		addx 16
		addx -11
		noop
		noop
		addx 21
		addx -15
		noop
		noop
		addx -3
		addx 9
		addx 1
		addx -3
		addx 8
		addx 1
		addx 5
		noop
		noop
		noop
		noop
		noop
		addx -36
		noop
		addx 1
		addx 7
		noop
		noop
		noop
		addx 2
		addx 6
		noop
		noop
		noop
		noop
		noop
		addx 1
		noop
		noop
		addx 7
		addx 1
		noop
		addx -13
		addx 13
		addx 7
		noop
		addx 1
		addx -33
		noop
		noop
		noop
		addx 2
		noop
		noop
		noop
		addx 8
		noop
		addx -1
		addx 2
		addx 1
		noop
		addx 17
		addx -9
		addx 1
		addx 1
		addx -3
		addx 11
		noop
		noop
		addx 1
		noop
		addx 1
		noop
		noop
		addx -13
		addx -19
		addx 1
		addx 3
		addx 26
		addx -30
		addx 12
		addx -1
		addx 3
		addx 1
		noop
		noop
		noop
		addx -9
		addx 18
		addx 1
		addx 2
		noop
		noop
		addx 9
		noop
		noop
		noop
		addx -1
		addx 2
		addx -37
		addx 1
		addx 3
		noop
		addx 15
		addx -21
		addx 22
		addx -6
		addx 1
		noop
		addx 2
		addx 1
		noop
		addx -10
		noop
		noop
		addx 20
		addx 1
		addx 2
		addx 2
		addx -6
		addx -11
		noop
		noop
		noop`}, 13140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drawScreen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{`addx 15
		addx -11
		addx 6
		addx -3
		addx 5
		addx -1
		addx -8
		addx 13
		addx 4
		noop
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx 5
		addx -1
		addx -35
		addx 1
		addx 24
		addx -19
		addx 1
		addx 16
		addx -11
		noop
		noop
		addx 21
		addx -15
		noop
		noop
		addx -3
		addx 9
		addx 1
		addx -3
		addx 8
		addx 1
		addx 5
		noop
		noop
		noop
		noop
		noop
		addx -36
		noop
		addx 1
		addx 7
		noop
		noop
		noop
		addx 2
		addx 6
		noop
		noop
		noop
		noop
		noop
		addx 1
		noop
		noop
		addx 7
		addx 1
		noop
		addx -13
		addx 13
		addx 7
		noop
		addx 1
		addx -33
		noop
		noop
		noop
		addx 2
		noop
		noop
		noop
		addx 8
		noop
		addx -1
		addx 2
		addx 1
		noop
		addx 17
		addx -9
		addx 1
		addx 1
		addx -3
		addx 11
		noop
		noop
		addx 1
		noop
		addx 1
		noop
		noop
		addx -13
		addx -19
		addx 1
		addx 3
		addx 26
		addx -30
		addx 12
		addx -1
		addx 3
		addx 1
		noop
		noop
		noop
		addx -9
		addx 18
		addx 1
		addx 2
		noop
		noop
		addx 9
		noop
		noop
		noop
		addx -1
		addx 2
		addx -37
		addx 1
		addx 3
		noop
		addx 15
		addx -21
		addx 22
		addx -6
		addx 1
		noop
		addx 2
		addx 1
		noop
		addx -10
		noop
		noop
		addx 20
		addx 1
		addx 2
		addx 2
		addx -6
		addx -11
		noop
		noop
		noop`}, `##..##..##..##..##..##..##..##..##..##..
		###...###...###...###...###...###...###.
		####....####....####....####....####....
		#####.....#####.....#####.....#####.....
		######......######......######......####
		#######.......#######.......#######.....
		`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawScreen(tt.args.s); got != tt.want {
				t.Errorf("drawScreen() = %v, want %v", got, tt.want)
			}
		})
	}
}
