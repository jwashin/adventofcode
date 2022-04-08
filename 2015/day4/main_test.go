package main

import "testing"

func TestFindMd5StartingWith00000(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{"abcdef"}, 609043},
		{"Example2", args{"pqrstuv"}, 1048970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMd5StartingWith00000(tt.args.aString); got != tt.want {
				t.Errorf("FindMd5StartingWith00000() = %v, want %v", got, tt.want)
			}
		})
	}
}
