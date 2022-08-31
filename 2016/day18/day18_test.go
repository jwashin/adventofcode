package main

import (
	"reflect"
	"testing"
)

func Test_trapRow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"..^^."}, ".^^^^"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRow(tt.args.s); got != tt.want {
				t.Errorf("trapRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeLines(t *testing.T) {
	type args struct {
		start string
		n     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"..^^.", 3}, []string{"..^^.", ".^^^^", "^^..^"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLines(tt.args.start, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSafeTiles(t *testing.T) {
	type args struct {
		start string
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{".^^.^.^^^^", 10}, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSafeTiles(tt.args.start, tt.args.n); got != tt.want {
				t.Errorf("countSafeTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeLines2(t *testing.T) {
	type args struct {
		start string
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{".^^.^.^^^^", 10}, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLines2(tt.args.start, tt.args.n); got != tt.want {
				t.Errorf("makeLines2() = %v, want %v", got, tt.want)
			}
		})
	}
}
