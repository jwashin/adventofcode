package main

import (
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test2(); got != tt.want {
				t.Errorf("test2() = %v, want %v", got, tt.want)
			}
		})
	}
}
