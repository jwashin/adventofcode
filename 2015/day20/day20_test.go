package main

import (
	"testing"
)

func Test_howManyPresents(t *testing.T) {
	type args struct {
		houseNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 10},
		{"2", args{2}, 30},
		{"3", args{3}, 40},
		{"4", args{4}, 70},
		{"5", args{5}, 60},
		{"6", args{6}, 120},
		{"7", args{7}, 80},
		{"8", args{8}, 150},
		{"9", args{9}, 130},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyPresents(tt.args.houseNumber); got != tt.want {
				t.Errorf("howManyPresents() = %v, want %v", got, tt.want)
			}
		})
	}
}
