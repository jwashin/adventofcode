package main

import "testing"

func Test_countTheWays(t *testing.T) {
	type args struct {
		s           string
		groupCounts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`???.###`, []int{1, 1, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTheWays(tt.args.s, tt.args.groupCounts); got != tt.want {
				t.Errorf("countTheWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
