package main

import (
	"testing"
)

func Test_location(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{"DR"}, 1, 1},
		{"2", args{"DUR"}, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := location(tt.args.path)
			if got != tt.want {
				t.Errorf("location() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("location() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findShortestPath(t *testing.T) {
	type args struct {
		passcode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"hijkl"}, ""},
		{"1", args{"ihgpwlah"}, "DDRRRD"},
		{"2", args{"kglvqrro"}, "DDUDRLRRUDRD"},
		{"3", args{"ulqzkmiv"}, "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestPath(tt.args.passcode); got != tt.want {
				t.Errorf("findShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLongestPathLength(t *testing.T) {
	type args struct {
		passcode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"0", args{"hijkl"}, 0},
		{"1", args{"ihgpwlah"}, 370},
		{"2", args{"kglvqrro"}, 492},
		{"3", args{"ulqzkmiv"}, 830},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestPathLength(tt.args.passcode); got != tt.want {
				t.Errorf("findLongestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
