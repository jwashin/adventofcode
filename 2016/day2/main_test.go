package main

import (
	"testing"
)

func Test_nextButton(t *testing.T) {
	type args struct {
		currentButton string
		instruction   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{currentButton: "5", instruction: "U"}, "2"},
		{"2", args{currentButton: "2", instruction: "L"}, "1"},
		{"3", args{currentButton: "1", instruction: "L"}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextButton(tt.args.currentButton, tt.args.instruction); got != tt.want {
				t.Errorf("nextButton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextCode(t *testing.T) {
	type args struct {
		start        string
		instructions string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{start: "5", instructions: "ULL"}, "1"},
		{"2", args{start: "1", instructions: "RRDDD"}, "9"},
		{"3", args{start: "9", instructions: "LURDL"}, "8"},
		{"4", args{start: "8", instructions: "UUUUD"}, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextCode(tt.args.start, tt.args.instructions); got != tt.want {
				t.Errorf("codeNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allCodes(t *testing.T) {
	type args struct {
		start string
		data  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{start: "5", data: `ULL
RRDDD
LURDL
UUUUD`}, "1985"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCodes(tt.args.start, tt.args.data); got != tt.want {
				t.Errorf("allCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allCodes2(t *testing.T) {
	type args struct {
		start string
		data  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{start: "5", data: `ULL
RRDDD
LURDL
UUUUD`}, "5DB3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCodes2(tt.args.start, tt.args.data); got != tt.want {
				t.Errorf("allCodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
