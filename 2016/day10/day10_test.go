package main

import (
	"reflect"
	"testing"
)

func Test_initBots(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want map[string]*Bot
	}{
		{"1", args{`value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`}, map[string]*Bot{
			"bot 2":    {purse: []int{5, 2}, lowDest: "bot 1", highDest: "bot 0"},
			"output 1": {purse: []int{}, lowDest: "", highDest: ""},
			"bot 1":    {purse: []int{3}, lowDest: "output 1", highDest: "bot 0"},
			"output 2": {purse: []int{}, lowDest: "", highDest: ""},
			"output 0": {purse: []int{}, lowDest: "", highDest: ""},
			"bot 0":    {purse: []int{}, lowDest: "output 2", highDest: "output 0"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initBots(tt.args.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initBots() = %v, want %v", got, tt.want)
			}
		})
	}
}

// {"1", args{"value 5 goes to bot 2"}, []string{"value 5", "bot 2"}},

// {"1", args{"bot 1 gives low to output 1 and high to bot 0"}, []string{"bot 1", "output 1", "bot 0"}},
// {"2", args{"bot 0 gives low to output 2 and high to output 0"}, []string{"bot 0", "output 2", "output 0"}},
// {"3", args{"bot 2 gives low to bot 1 and high to bot 0"}, []string{"bot 2", "bot 1", "bot 0"}},

func Test_getValue(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{"1", args{"value 5 goes to bot 2"}, 5, "bot 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getValue(tt.args.s)
			if got != tt.want {
				t.Errorf("getValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getInstructions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{"1", args{"bot 1 gives low to output 1 and high to bot 0"}, "bot 1", "output 1", "bot 0"},
		{"2", args{"bot 0 gives low to output 2 and high to output 0"}, "bot 0", "output 2", "output 0"},
		{"3", args{"bot 2 gives low to bot 1 and high to bot 0"}, "bot 2", "bot 1", "bot 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := getInstructions(tt.args.s)
			if got != tt.want {
				t.Errorf("getInstructions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getInstructions() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("getInstructions() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_doInstructions(t *testing.T) {
	type args struct {
		bots map[string]*Bot
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{map[string]*Bot{
			"bot 2":    {purse: []int{5, 2}, lowDest: "bot 1", highDest: "bot 0"},
			"output 1": {purse: []int{}, lowDest: "", highDest: ""},
			"bot 1":    {purse: []int{3}, lowDest: "output 1", highDest: "bot 0"},
			"output 2": {purse: []int{}, lowDest: "", highDest: ""},
			"output 0": {purse: []int{}, lowDest: "", highDest: ""},
			"bot 0":    {purse: []int{}, lowDest: "output 2", highDest: "output 0"}}}, 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doInstructions(tt.args.bots); got != tt.want {
				t.Errorf("doInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
