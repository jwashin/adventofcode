package main

import (
	"testing"
)

func Test_calcMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc18"}, "0034e0923cc38887a57bd7b1d4f953df"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMD5(tt.args.s); got != tt.want {
				t.Errorf("calcMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstTriplet(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc18"}, "8"},
		{"2", args{"abc17"}, ""},
		{"3", args{"abc534"}, "7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstTriplet(tt.args.salt); got != tt.want {
				t.Errorf("findFirstTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makePad(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 22728},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makePad(tt.args.salt); got != tt.want {
				t.Errorf("makePad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc0"}, "a107ff634856bb300138cac6568c0f24"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superHash(tt.args.s); got != tt.want {
				t.Errorf("superHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSuperFirstTriplet(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc5"}, "2"},
		{"2", args{"abc10"}, "e"},
		// {"3", args{"abc534"}, "7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSuperFirstTriplet(tt.args.aString); got != tt.want {
				t.Errorf("findSuperFirstTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superMakePad(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 22551},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superMakePad(tt.args.salt); got != tt.want {
				t.Errorf("superMakePad() = %v, want %v", got, tt.want)
			}
		})
	}
}
