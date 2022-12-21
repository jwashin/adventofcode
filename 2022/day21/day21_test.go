package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`root: pppw + sjmn
		dbpl: 5
		cczh: sllz + lgvd
		zczc: 2
		ptdq: humn - dvpt
		dvpt: 3
		lfqf: 4
		humn: 5
		ljgn: 2
		sjmn: drzm * dbpl
		sllz: 4
		pppw: cczh / lfqf
		lgvd: ljgn * ptdq
		drzm: hmdt - zczc
		hmdt: 32`}, 152},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`root: pppw + sjmn
		dbpl: 5
		cczh: sllz + lgvd
		zczc: 2
		ptdq: humn - dvpt
		dvpt: 3
		lfqf: 4
		humn: 5
		ljgn: 2
		sjmn: drzm * dbpl
		sllz: 4
		pppw: cczh / lfqf
		lgvd: ljgn * ptdq
		drzm: hmdt - zczc
		hmdt: 32`}, 301},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
