package main

import (
	"reflect"
	"testing"
)

func Test_singleReplacement(t *testing.T) {
	type args struct {
		s    string
		from string
		to   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"HOH", "H", "HO"}, []string{"HOOH", "HOHO"}},
		{"3", args{"HOH", "O", "HH"}, []string{"HHHH"}},
		{"2", args{"HOH", "H", "OH"}, []string{"OHOH", "HOOH"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleReplacement(tt.args.s, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moleculeCount(t *testing.T) {
	type args struct {
		s     string
		rules []rule
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1.", args{"HOH", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}}}, 4},
		{"2.", args{"HOHOHO", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}}}, 7},
		{"3.", args{"ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF", []rule{{"B", "BCa"}}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moleculeCount(tt.args.s, tt.args.rules); got != tt.want {
				t.Errorf("moleculeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fewestSteps(t *testing.T) {
	type args struct {
		target string
		rules  []rule
		start  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"HOH", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}, {"e", "H"}, {"e", "O"}}, "e"}, 3},
		{"2", args{"HOHOHO", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}, {"e", "H"}, {"e", "O"}}, "e"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fewestSteps(tt.args.target, tt.args.rules, tt.args.start); got != tt.want {
				t.Errorf("fewestSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleUnReplacement(t *testing.T) {
	type args struct {
		s          string
		toString   string
		fromString string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"HOOH", "HO", "H"}, []string{"HOH"}},
		{"2", args{"HOHO", "HO", "H"}, []string{"HOH", "HHO"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleUnReplacement(tt.args.s, tt.args.toString, tt.args.fromString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleUnReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fewestStepsA(t *testing.T) {
	type args struct {
		target string
		rules  []rule
		start  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"e", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}, {"e", "H"}, {"e", "O"}}, "HOH"}, 3},
		{"2", args{"e", []rule{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}, {"e", "H"}, {"e", "O"}}, "HOHOHO"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fewestStepsA(tt.args.target, tt.args.rules, tt.args.start); got != tt.want {
				t.Errorf("fewestStepsA() = %v, want %v", got, tt.want)
			}
		})
	}
}
