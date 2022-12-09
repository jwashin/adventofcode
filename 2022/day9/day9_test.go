package main

import "testing"

func Test_follow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`R 4
		U 4
		L 3
		D 1
		R 4
		D 1
		L 5
		R 2`}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := follow(tt.args.s); got != tt.want {
				t.Errorf("follow() = %v, want %v", got, tt.want)
			}
		})
	}
}
