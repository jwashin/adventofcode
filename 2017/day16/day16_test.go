package main

import (
	"testing"
)

func Test_spin(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcde", 3}, "cdeab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spin(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("spin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exchange(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"eabcd", 3, 4}, "eabdc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partner(t *testing.T) {
	type args struct {
		s string
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"eabdc", "e", "b"}, "baedc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partner(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("partner() = %v, want %v", got, tt.want)
			}
		})
	}
}
