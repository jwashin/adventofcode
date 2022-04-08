package main

import (
	"reflect"
	"testing"
)

func TestCountEasyDigits(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
		edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
		fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
		fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
		aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
		fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
		dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
		bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
		egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
		gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountEasyDigits(tt.args.aString); got != tt.want {
				t.Errorf("CountEasyDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortItems(t *testing.T) {
	type args struct {
		aString []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"one", args{[]string{"febca", "cfagb", "ecbafd", "efdcbg", "cbegdfa", "fg", "bgafec", "gfae", "acgdb", "gfc"}}, []string{"abcef", "abcfg", "abcdef", "bcdefg", "abcdefg", "fg", "abcefg", "aefg", "abcdg", "cfg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortItems(tt.args.aString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_makeCode(t *testing.T) {
// 	type args struct {
// 		oldVal []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{"Example", args{[]string{
// 			"acedgfb",
// 			"cdfbe",
// 			"gcdfa",
// 			"fbcad",
// 			"dab",
// 			"cefabd",
// 			"cdfgeb",
// 			"eafb",
// 			"cagedb",
// 			"ab",
// 		}}, "deafgbc"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := makeCode(tt.args.oldVal); got != tt.want {
// 				t.Errorf("makeCode() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_lineValue(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}, 5353},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lineValue(tt.args.line); got != tt.want {
				t.Errorf("lineValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateOutputTotals(t *testing.T) {
	type args struct {
		bigString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
		edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
		fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
		fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
		aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
		fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
		dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
		bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
		egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
		gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`}, 61229},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateOutputTotals(tt.args.bigString); got != tt.want {
				t.Errorf("CalculateOutputTotals() = %v, want %v", got, tt.want)
			}
		})
	}
}
