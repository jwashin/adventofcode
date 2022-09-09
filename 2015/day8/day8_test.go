package main

import (
	"testing"
)

func Test_strCounts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{`""`}, 2, 0},
		{"2", args{`"ABC"`}, 5, 3},
		{"3", args{`"\\"`}, 4, 1},
		{"4", args{`"\""`}, 4, 1},
		{"5", args{`"aaa\"aaa"`}, 10, 7},
		{"6", args{`"\x27"`}, 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := strCounts(tt.args.s)
			if got != tt.want {
				t.Errorf("strCounts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("strCounts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_strunCounts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{`""`}, 2, 6},
		{"2", args{`"abc"`}, 5, 9},
		{"5", args{`"aaa\"aaa"`}, 10, 16},
		{"3", args{`"\\"`}, 4, 10},
		{"4", args{`"\""`}, 4, 10},
		{"6", args{`"\x27"`}, 6, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := strunCounts(tt.args.s)
			if got != tt.want {
				t.Errorf("strunCounts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("strunCounts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
