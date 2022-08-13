package main

import (
	"reflect"
	"testing"
)

func Test_parseData(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want []Scanner
	}{
		{"1", args{`--- scanner 0 ---
		0,2
		4,1
		3,3
		
		--- scanner 1 ---
		-1,-1
		-5,0
		-2,1`}, []Scanner{
			{0, []Vector{{0, 2, 0}, {4, 1, 0}, {3, 3, 0}}, Vector{0, 0, 0}, 0},
			{1, []Vector{{-1, -1, 0}, {-5, 0, 0}, {-2, 1, 0}}, Vector{0, 0, 0}, 0}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseData(tt.args.aString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
