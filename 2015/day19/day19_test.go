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
