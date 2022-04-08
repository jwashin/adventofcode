package main

import (
	"testing"
)

// {"1", args{`#############
// #...........#
// ###B#C#B#D###
//   #A#D#C#A#
//   #########`, `#############
//   #...........#
//   ###A#B#C#D###
// 	#A#B#C#D#
// 	#########`}, 12521}, // TODO: Add test cases.

func TestTableau_doMove(t *testing.T) {
	type fields struct {
		data   []string
		cost   int
		parent *Tableau
	}
	type args struct {
		m Move
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"1", fields{[]string{"#############", "#...........#", "###B#C#B#D###",
			"  #A#D#C#A#  ",
			"  #########  "}, 0, nil}, args{Move{7, 4}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tableau{
				data:   tt.fields.data,
				cost:   tt.fields.cost,
				parent: tt.fields.parent,
			}
			if got := tr.doMove(tt.args.m); got != tt.want {
				t.Errorf("Tableau.doMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinCost(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`#############
#...........#
###B#C#B#D###
#D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`}, 44169},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinCost(tt.args.aString); got != tt.want {
				t.Errorf("findMinCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
