package main

import "testing"

// 		{"1", args{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}, 3068},

func Test_doTetris(t *testing.T) {
	type args struct {
		s     string
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1", args{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>", 2022}, 3068},
		{"2", args{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>", 1000000000000}, 1514285714288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doTetris(tt.args.s, tt.args.count); got != tt.want {
				t.Errorf("doTetris() = %v, want %v", got, tt.want)
			}
		})
	}
}
