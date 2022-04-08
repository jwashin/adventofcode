package elevator

import (
	"testing"
)

func TestGetFloor(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0a", args{"(())"}, 0},
		{"0b", args{"()()"}, 0},
		{"3a", args{"((("}, 3},
		{"3b", args{"(()(()("}, 3},
		{"3c", args{"))((((("}, 3},
		{"-1a", args{"())"}, -1},
		{"-1b", args{"))("}, -1},
		{"-3a", args{")))"}, -3},
		{"-3b", args{")())())"}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloor(tt.args.instructions); got != tt.want {
				t.Errorf("GetFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPositionX(t *testing.T) {
	type args struct {
		instructions string
		x            int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2a", args{")", -1}, 1},
		{"2b", args{"()())", -1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPositionX(tt.args.instructions, tt.args.x); got != tt.want {
				t.Errorf("GetPositionX() = %v, want %v", got, tt.want)
			}
		})
	}
}
