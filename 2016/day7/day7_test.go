package main

import (
	"reflect"
	"testing"
)

func Test_bracketParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want BracketParsed
	}{
		{"1", args{"ioxxoj[asdfgh]zxcvbn"}, BracketParsed{inside: []string{"asdfgh"}, outside: []string{"ioxxoj", "zxcvbn"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bracketParse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bracketParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_has4palindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abba"}, true},
		{"2", args{"mnop"}, false},
		{"3", args{"ioxxo"}, true},
		{"4", args{"aaaa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := has4palindrome(tt.args.s); got != tt.want {
				t.Errorf("has4palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_supportsTLS(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abba[mnop]qrst"}, true},
		{"2", args{"abcd[bddb]xyyx"}, false},
		{"3", args{"aaaa[qwer]tyui"}, false},
		{"4", args{"ioxxoj[asdfgh]zxcvbn"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supportsTLS(tt.args.s); got != tt.want {
				t.Errorf("supportsTLS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getABAs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"aba"}, []string{"aba"}},
		{"2", args{"zazbz"}, []string{"zaz", "zbz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getABAs(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getABAs() = %v, want %v", got, tt.want)
			}
		})
	}
}
