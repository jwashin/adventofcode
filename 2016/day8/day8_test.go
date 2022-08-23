package main

import (
	"reflect"
	"testing"
)

func Test_rotateString(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"##.", 1}, ".##"},
		{"2", args{"#.#....", 4}, "....#.#"},
		{"3", args{".##", 1}, "#.#"},
		{"5", args{"01234", 1}, "40123"},
		{"4", args{"01234", 2}, "34012"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rect(t *testing.T) {
	type args struct {
		r []string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{".......", ".......", "......."}, 3, 2}, []string{"###....", "###....", "......."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rect(tt.args.r, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotCol(t *testing.T) {
	type args struct {
		r []string
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"###....", "###....", "......."}, 1, 1}, []string{"#.#....", "###....", ".#....."}},
		{"2", args{[]string{"....#.#", "###....", ".#....."}, 1, 1}, []string{".#..#.#", "#.#....", ".#....."}},
		{"3", args{[]string{"00000", "11111", "22222"}, 2, 1}, []string{"00200", "11011", "22122"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotCol(tt.args.r, tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotCol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotRow(t *testing.T) {
	type args struct {
		r []string
		y int
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"#.#....", "###....", ".#....."}, 0, 4}, []string{"....#.#", "###....", ".#....."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotRow(tt.args.r, tt.args.y, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
