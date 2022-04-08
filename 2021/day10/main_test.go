package main

import (
	"testing"
)

func Test_findFirstIllegalCharacter(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"{([(<{}[<>[]}>{[]{[(<()>"}, "}"},
		{"", args{"[[<[([]))<([[{}[[()]]]"}, ")"},
		{"", args{"[{[{({}]{}}([{[{{{}}([]"}, "]"},
		{"", args{"[<(<(<(<{}))><([]([]()"}, ")"},
		{"", args{"<{([([[(<>()){}]>(<<{{"}, ">"},
		{"", args{"[({(<(())[]>[[{[]{<()<>>"}, "}}]])})]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstIllegalCharacter(tt.args.aString); got != tt.want {
				t.Errorf("findFirstIllegalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreIllegals(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`[({(<(())[]>[[{[]{<()<>>
			[(()[<>])]({[<{<<[]>>(
			{([(<{}[<>[]}>{[]{[(<()>
			(((({<>}<{<{<>}{[]{[]{}
			[[<[([]))<([[{}[[()]]]
			[{[{({}]{}}([{[{{{}}([]
			{<[[]]>}<{[{[{[]{()[[[]
			[<(<(<(<{}))><([]([]()
			<{([([[(<>()){}]>(<<{{
			<{([{{}}[<[[[<>{}]]]>[]]`}, 26397},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreIllegals(tt.args.aString); got != tt.want {
				t.Errorf("scoreIllegals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreIncompletes(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`[({(<(())[]>[[{[]{<()<>>
			[(()[<>])]({[<{<<[]>>(
			{([(<{}[<>[]}>{[]{[(<()>
			(((({<>}<{<{<>}{[]{[]{}
			[[<[([]))<([[{}[[()]]]
			[{[{({}]{}}([{[{{{}}([]
			{<[[]]>}<{[{[{[]{()[[[]
			[<(<(<(<{}))><([]([]()
			<{([([[(<>()){}]>(<<{{
			<{([{{}}[<[[[<>{}]]]>[]]`}, 288957},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreIncompletes(tt.args.aString); got != tt.want {
				t.Errorf("scoreIncompletes() = %v, want %v", got, tt.want)
			}
		})
	}
}
