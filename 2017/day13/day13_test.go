package main

import (
	"testing"
)

func Test_severity(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0: 3
		1: 2
		4: 4
		6: 4`}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := severity(tt.args.s); got != tt.want {
				t.Errorf("severity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcDelay(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0: 3
		1: 2
		4: 4
		6: 4`}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDelay(tt.args.s); got != tt.want {
				t.Errorf("calcDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}
