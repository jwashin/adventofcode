package main

import (
	"reflect"
	"testing"
)

func Test_sortFloor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"E,HM,LM,HG,LG"}, "E,HG,HM,LG,LM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortFloor(tt.args.s); got != tt.want {
				t.Errorf("sortFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinations2(t *testing.T) {
	type args struct {
		set []string
		n   int
	}
	tests := []struct {
		name        string
		args        args
		wantSubsets [][]string
	}{
		{"1", args{[]string{"HG", "HM", "LG"}, 2}, [][]string{{"HG", "HM"}, {"HG", "LG"}, {"HM", "LG"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSubsets := Combinations2(tt.args.set, tt.args.n); !reflect.DeepEqual(gotSubsets, tt.wantSubsets) {
				t.Errorf("Combinations2() = %v, want %v", gotSubsets, tt.wantSubsets)
			}
		})
	}
}

func Test_isSafeFloor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{"PG,RG,RM,WG,WM,XG"}, true},
		{"1", args{"E,HG,LG,LM"}, true},
		{"2", args{"E,HG,LG"}, true},
		{"3", args{"LM"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafeFloor(tt.args.s); got != tt.want {
				t.Errorf("isSafeFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
