package main

import "testing"

func Test_decompress(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"ADVENT"}, "ADVENT"},
		{"2", args{"A(1x5)BC"}, "ABBBBBC"},
		{"3", args{"(3x3)XYZ"}, "XYZXYZXYZ"},
		{"4", args{"A(2x2)BCD(2x2)EFG"}, "ABCBCDEFEFG"},
		{"5", args{"(6x1)(1x3)A"}, "(1x3)A"},
		{"6", args{"X(8x2)(3x3)ABCY"}, "X(3x3)ABC(3x3)ABCY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompress(tt.args.s); got != tt.want {
				t.Errorf("decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
