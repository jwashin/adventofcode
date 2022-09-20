package main

import (
	"testing"
)

func Test_countMatchesNext(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1122"}, 3},
		{"2", args{"1111"}, 4},
		{"3", args{"1234"}, 0},
		{"4", args{"91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchesNext(tt.args.s); got != tt.want {
				t.Errorf("countMatchesNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMatchesHalfway(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1212"}, 6},
		{"2", args{"1221"}, 0},
		{"3", args{"123425"}, 4},
		{"4", args{"123123"}, 12},
		{"5", args{"12131415"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchesHalfway(tt.args.s); got != tt.want {
				t.Errorf("countMatchesHalfway() = %v, want %v", got, tt.want)
			}
		})
	}
}
