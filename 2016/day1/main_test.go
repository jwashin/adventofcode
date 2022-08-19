package main

import (
	"reflect"
	"testing"
)

func Test_parseDirection(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want vector
	}{
		{"1", args{"R2"}, vector{direction: "R", distance: 2}},
		{"2", args{"L3"}, vector{direction: "L", distance: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDirections(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_getDistance(t *testing.T) {
// 	type args struct {
// 		input string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{"1", args{"R2, L3"}, 5},
// 		{"2", args{"R2, R2, R2"}, 2},
// 		{"3", args{"R5, L5, R5, R3"}, 12},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := getDistance(tt.args.input); got != tt.want {
// 				t.Errorf("getDistance() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_getDistance(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{"R8, R4, R4, R8"}, 8, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getDistance(tt.args.input)
			if got != tt.want {
				t.Errorf("getDistance() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getDistance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
