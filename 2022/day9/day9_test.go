package main

import "testing"

// {"1", args{`R 4
// U 4
// L 3
// D 1
// R 4
// D 1
// L 5
// R 2`}, 13},

func Test_follow(t *testing.T) {
	type args struct {
		s     string
		count int
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
R 2`, 2}, 13},
		{"2", args{`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`, 10}, 1},
		{"3", args{`R 5
		U 8
		L 8
		D 3
		R 17
		D 10
		L 25
		U 20`, 10}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := follow(tt.args.s, tt.args.count); got != tt.want {
				t.Errorf("follow() = %v, want %v", got, tt.want)
			}
		})
	}
}
