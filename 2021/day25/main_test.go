package main

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		aTableau []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"123", "456", "789"}}, []string{"147", "258", "369"}},
		{"2", args{[]string{"1234", "4567", "7890"}}, []string{"147", "258", "369", "470"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.aTableau); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doStep(t *testing.T) {
	type args struct {
		aTableau []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{"0", args{[]string{
			"..........",
			".>v....v..",
			".......>..",
			"..........",
		}}, []string{
			"..........",
			".>........",
			"..v....v>.",
			"..........",
		}},

		{"1", args{[]string{"...>...",
			".......",
			"......>",
			"v.....>",
			"......>",
			".......",
			"..vvv.."}}, []string{
			"..vv>..",
			".......",
			">......",
			"v.....>",
			">......",
			".......",
			"....v.."}},

		{"2", args{[]string{
			"..vv>..",
			".......",
			">......",
			"v.....>",
			">......",
			".......",
			"....v.."}}, []string{
			"....v>.",
			"..vv...",
			".>.....",
			"......>",
			"v>.....",
			".......",
			"......."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doStep(tt.args.aTableau); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveRight(t *testing.T) {
	type args struct {
		tableau []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"...>>>>>..."}}, []string{"...>>>>.>.."}},
		{"2", args{[]string{"...>>>>.>.."}}, []string{"...>>>.>.>."}},
		{"3", args{[]string{"...>>>.>.>."}}, []string{"...>>.>.>.>"}},
		{"4", args{[]string{"...>>.>.>.>"}}, []string{">..>.>.>.>."}},
		{"5", args{[]string{">..>.>.>.>."}}, []string{".>..>.>.>.>"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveRight(tt.args.tableau); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stepsToStop(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`v...>>.vv>
		.vv>>.vv..
		>>.>v>...v
		>>v>>.>.v.
		v>v.vv.v..
		>.>>..v...
		.vv..>.>v.
		v.v..>>v.v
		....v..v.>`}, 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepsToStop(tt.args.aString); got != tt.want {
				t.Errorf("stepsToStop() = %v, want %v", got, tt.want)
			}
		})
	}
}
