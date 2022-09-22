package main

import (
	"testing"
)

// {"<>,", args{"<>"}, true},
// 		{"<random characters>,", args{"<random characters>"}, true},
// 		{"<<<<>,", args{"<<<<>"}, true},
// 		{"<{!>}>,", args{"<{!>}>"}, true},
// 		{"<!!>,", args{"<!!>"}, true},
// 		{"3", args{"<!>,>"}, true},
// 		{"<!!!>>,", args{"<!!!>>"}, true},
// 		{"<{o\"i!a,<{i<a>,", args{"<{o\"i!a,<{i<a>"}, true},
// 		{"9", args{"<a!>},{<a!>},{<a!>},{<ab>"}, true},
// 		{"9", args{"<a!>"}, false},

func Test_removeGarbage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"<>"}, ""},
		{"2", args{"<random characters>"}, ""},
		{"3", args{"<<<<>"}, ""},
		{"4", args{"<{!>}>"}, ""},
		{"5", args{"<!!>"}, ""},
		{"6", args{"<!!!>>"}, ""},
		{"7", args{"<{o\"i!a,<{i<a>"}, ""},
		{"6", args{"<!>,>"}, ""},
		{"9", args{"<a!>},{<a!>},{<a!>},{<ab>"}, ""},
		{"z", args{"<a!>abc>"}, ""},
		{"t", args{"{{<!>},{<!>},{<!>},{<a>}}"}, "{{}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeGarbage(tt.args.s); got != tt.want {
				t.Errorf("removeGarbage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreGroups(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"{}"}, 1},
		{"2", args{"{{{}}}"}, 6},
		{"3", args{"{{},{}}"}, 5},
		{"4", args{"{{{},{},{{}}}}"}, 16},
		{"5", args{"{<a>,<a>,<a>,<a>}"}, 1},
		{"6", args{"{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9},
		{"7", args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9},
		{"8", args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreGroups(tt.args.s); got != tt.want {
				t.Errorf("scoreGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countGarbage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"<>"}, 0},
		{"2", args{"<random characters>"}, 17},
		{"3", args{"<<<<>"}, 3},
		{"4", args{"<{!>}>"}, 2},
		{"5", args{"<!!>"}, 0},
		{"6", args{"<!!!>>"}, 0},
		{"7", args{"<{o\"i!a,<{i<a>"}, 10},
		{"8", args{"<!>,>"}, 1},
		{"9", args{"<a!>},{<a!>},{<a!>},{<ab>"}, 17},
		{"z", args{"<a!>abc>"}, 4},
		{"t", args{"{{<!>},{<!>},{<!>},{<a>}}"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGarbage(tt.args.s); got != tt.want {
				t.Errorf("countGarbage() = %v, want %v", got, tt.want)
			}
		})
	}
}
