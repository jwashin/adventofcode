package main

import (
	"testing"
)

func Test_test17(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test17(); got != tt.want {
				t.Errorf("test17() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testa17(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testa17(); got != tt.want {
				t.Errorf("testa17() = %v, want %v", got, tt.want)
			}
		})
	}
}
