package main

import (
	"reflect"
	"testing"
)

func Test_isATriangle(t *testing.T) {
	type args struct {
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"5", "10", "25"}, false},
		{"2", args{"30", "30", "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleTriangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("isATriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{" 775  785  361"}, []string{"775", "785", "361"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
