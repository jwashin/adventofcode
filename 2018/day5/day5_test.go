package main

import "testing"

func Test_react(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"dabAcCaCBAcCcaDA"}, "dabCBAcaDA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := react(tt.args.s); got != tt.want {
				t.Errorf("react() = %v, want %v", got, tt.want)
			}
		})
	}
}
