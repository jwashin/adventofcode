package main

import (
	"testing"
)

func Test_countItems(t *testing.T) {
	type args struct {
		elementCount int
		target       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{2, 100}, 99},
		{"2.", args{4, 100}, 156849},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countItems(tt.args.elementCount, tt.args.target); got != tt.want {
				t.Errorf("countItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximizeScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`}, 62842880},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeScore(tt.args.s); got != tt.want {
				t.Errorf("maximizeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximizeScore500(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`}, 57600000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeScore500(tt.args.s); got != tt.want {
				t.Errorf("maximizeScore500() = %v, want %v", got, tt.want)
			}
		})
	}
}
