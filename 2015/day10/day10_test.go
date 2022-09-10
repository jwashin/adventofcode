package main

import "testing"

func Test_lookandsay(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1"}, "11"},
		{"2", args{"11"}, "21"},
		{"3", args{"21"}, "1211"},
		{"4", args{"1211"}, "111221"},
		{"5", args{"111221"}, "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookandsay(tt.args.s); got != tt.want {
				t.Errorf("lookandsay() = %v, want %v", got, tt.want)
			}
		})
	}
}
