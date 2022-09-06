package main

import "testing"

func Test_test(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`}, 3},

		{"2", args{`cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.s); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
