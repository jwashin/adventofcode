package main

import (
	"reflect"
	"testing"
)

func Test_winningElf(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2", args{2}, 1},
		{"3", args{3}, 3},
		{"4", args{4}, 1},
		{"5", args{5}, 3},
		{"6", args{6}, 5},
		{"7", args{7}, 7},
		{"8", args{8}, 1},
		{"9", args{9}, 3},
		{"10", args{10}, 5},
		{"11", args{11}, 7},
		{"12", args{12}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winningElf(tt.args.n); got != tt.want {
				t.Errorf("winningElf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winningElf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5", args{5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winningElf2(tt.args.n); got != tt.want {
				t.Errorf("winningElf2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winningElf3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5", args{5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winningElf3(tt.args.n); got != tt.want {
				t.Errorf("winningElf3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeOne2(t *testing.T) {
	type args struct {
		currentChooser int
		elves          []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"1", args{0, []int{1, 2, 3, 4, 5}}, 1, []int{1, 2, 4, 5}},
		{"2", args{1, []int{1, 2, 4, 5}}, 2, []int{1, 2, 4}},
		{"3", args{2, []int{1, 2, 4}}, 0, []int{2, 4}},
		{"4", args{0, []int{2, 4}}, 1, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeOne2(tt.args.currentChooser, tt.args.elves)
			if got != tt.want {
				t.Errorf("removeOne2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeOne2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
