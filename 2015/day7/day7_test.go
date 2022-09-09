package main

import "testing"

func Test_test(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"1", args{"d"}, 72},
		{"2", args{"e"}, 507},
		{"3", args{"f"}, 492},
		{"4", args{"g"}, 114},
		{"5", args{"h"}, 65412},
		{"6", args{"i"}, 65079},
		{"7", args{"x"}, 123},
		{"8", args{"y"}, 456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.s); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
