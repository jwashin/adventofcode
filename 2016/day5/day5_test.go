package main

import (
	"testing"
)

func Test_findPassword(t *testing.T) {
	type args struct {
		doorId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc"}, "18f47a30"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPassword(tt.args.doorId); got != tt.want {
				t.Errorf("findPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc3231929"}, "00000155f8105dff7f56ee10fa9b9abd"},
		{"2", args{"abc5017308"}, "000008f82c5b3924a1ecbebf60344e00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.s); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPassword2(t *testing.T) {
	type args struct {
		doorId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc"}, "05ace8e3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPassword2(tt.args.doorId); got != tt.want {
				t.Errorf("findPassword2() = %v, want %v", got, tt.want)
			}
		})
	}
}
