package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		passphrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aa bb cc dd ee"}, true},
		{"2", args{"aa bb cc dd aa \n"}, false},
		{"3", args{"aa bb  cc dd aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.passphrase); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
