package main

import (
	"testing"
)

// {"1", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, 7},
// 		{"2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, 5},
// 		{"3", args{"nppdvjthqldpwncqszvftbrmjlhg"}, 6},
// 		{"4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, 10},
// 		{"5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, 11},

func Test_startOfPacket(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4}, 7},
		{"2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4}, 5},
		{"3", args{"nppdvjthqldpwncqszvftbrmjlhg", 4}, 6},
		{"4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4}, 10},
		{"5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4}, 11},

		{"1a", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14}, 19},
		{"2a", args{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14}, 23},
		{"3a", args{"nppdvjthqldpwncqszvftbrmjlhg", 14}, 23},
		{"4a", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14}, 29},
		{"5a", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfPacket(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("startOfPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}
