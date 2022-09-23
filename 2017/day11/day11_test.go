package main

import "testing"

func Test_hexStepsAway(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ne,ne,ne"}, 3},
		{"2", args{"ne,ne,sw,sw"}, 0},
		{"3", args{"ne,ne,s,s"}, 2},
		{"4", args{"se,sw,se,sw,sw"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexStepsAway(tt.args.data); got != tt.want {
				t.Errorf("hexStepsAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
