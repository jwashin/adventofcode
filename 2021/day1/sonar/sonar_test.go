package sonar

import (
	"testing"
)

func TestCountIncreases(t *testing.T) {
	type args struct {
		aList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]string{"199",
			"200",
			"208",
			"210",
			"200",
			"207",
			"240",
			"269",
			"260",
			"263"}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIncreases(tt.args.aList); got != tt.want {
				t.Errorf("CountIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountIncreasesbyThrees(t *testing.T) {
	type args struct {
		aList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]string{"199",
			"200",
			"208",
			"210",
			"200",
			"207",
			"240",
			"269",
			"260",
			"263"}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIncreasesbyThrees(tt.args.aList); got != tt.want {
				t.Errorf("CountIncreasesbyThrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
