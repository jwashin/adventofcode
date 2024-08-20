package main

import "testing"

func Test_countArrangements(t *testing.T) {
	type args struct {
		s      string
		counts string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"???.###", "1,1,3"}, 1},
		{"2", args{".??..??...?##.", "1,1,3"}, 4},
		{"3", args{"?#?#?#?#?#?#?#?", "1,3,1,6"}, 1},
		{"4", args{"????.#...#...", "4,1,1"}, 1},
		{"5", args{"????.######..#####.", "1,6,5"}, 4},
		{"6", args{"?###????????", "3,2,1"}, 10},
		{"6a", args{"???????", "2,1"}, 10},
		{"6b", args{"??????", "2,1"}, 6},
		{"6c", args{"?????", "2,1"}, 3},
		{"6d", args{"????", "2,1"}, 1},
		{"6e", args{"???", "2,1"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangements(tt.args.s, tt.args.counts); got != tt.want {
				t.Errorf("countArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
