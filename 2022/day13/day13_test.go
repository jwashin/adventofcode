package main

import (
	"reflect"
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
		{"1", args{`[1,1,3,1,1]
		[1,1,5,1,1]
		
		[[1],[2,3,4]]
		[[1],4]
		
		[9]
		[[8,7,6]]
		
		[[4,4],4,4]
		[[4,4],4,4,4]
		
		[7,7,7,7]
		[7,7,7]
		
		[]
		[3]
		
		[[[]]]
		[[]]
		
		[1,[2,[3,[4,[5,6,7]]]],8,9]
		[1,[2,[3,[4,[5,6,0]]]],8,9]`}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseList(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"[1,1,3,1,1,"}, []string{"1", "1", "3", "1", "1"}},
		{"2", args{"[[1],[2,3,4]]"}, []string{"[1]", "[2,3,4]"}},
		{"3", args{"[[1],4]"}, []string{"[1]", "4"}},
		{"4", args{"[[]]"}, []string{"[]"}},
		{"5", args{"[[[]]]"}, []string{"[[]]"}},
		{"6", args{"[1,[2,[3,[4,[5,6,0]]]],8,9]"}, []string{"1", "[2,[3,[4,[5,6,0]]]]", "8", "9"}},
		{"7", args{"[]"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseList(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseList() = %v, want %v", got, tt.want)
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
		{"1", args{`[1,1,3,1,1]
		[1,1,5,1,1]
		
		[[1],[2,3,4]]
		[[1],4]
		
		[9]
		[[8,7,6]]
		
		[[4,4],4,4]
		[[4,4],4,4,4]
		
		[7,7,7,7]
		[7,7,7]
		
		[]
		[3]
		
		[[[]]]
		[[]]
		
		[1,[2,[3,[4,[5,6,7]]]],8,9]
		[1,[2,[3,[4,[5,6,0]]]],8,9]`}, 140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
