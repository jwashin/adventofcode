package main

import (
	"reflect"
	"testing"
)

func Test_candidateOpcodes(t *testing.T) {
	type args struct {
		before []int
		a      int
		b      int
		c      int
		after  []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{before: []int{3, 2, 1, 1}, a: 2, b: 1, c: 2, after: []int{3, 2, 2, 1}}, []string{"addi", "mulr", "seti"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candidateOpcodes(tt.args.before, tt.args.a, tt.args.b, tt.args.c, tt.args.after); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("candidateOpcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
