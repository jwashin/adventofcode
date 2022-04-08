package main

import (
	"reflect"
	"testing"
)

func Test_parsePair(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{"1", args{"[1,2]"}, "1", "2"},
		{"2", args{"[[1,2],3]"}, "[1,2]", "3"},
		{"3", args{"[9,[8,7]]"}, "9", "[8,7]"},
		{"4", args{"[[1,9],[8,5]]"}, "[1,9]", "[8,5]"},
		{"5", args{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"}, "[[[1,2],[3,4]],[[5,6],[7,8]]]", "9"},
		{"6", args{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"}, "[[9,[3,8]],[[0,9],6]]", "[[[3,7],[4,9]],3]"},
		{"7", args{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"}, "[[[1,3],[5,3]],[[1,3],[8,7]]]", "[[[4,9],[6,9]],[[8,2],[7,3]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parsePair(tt.args.aString)
			if got != tt.want {
				t.Errorf("parsePair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parsePair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcde"}, "edcba"},
		{"2", args{"ab"}, "ba"},
		{"3", args{"a"}, "a"},
		{"4", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.aString); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addToLastNumber(t *testing.T) {
	type args struct {
		aString string
		x       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"[[[[", 9}, "[[[["},
		{"2", args{"[7,[6,[5,[4,", 3}, "[7,[6,[5,[7,"},
		{"3", args{"[7,[6,[5,[0,", 3}, "[7,[6,[5,[3,"},
		{"4", args{"[[[[4,0],", 5}, "[[[[4,5],"},
		{"5", args{"[[[[5,9],[16,0]],[[10,", 1}, "[[[[5,9],[16,0]],[[11,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToLastNumber(tt.args.aString, tt.args.x); got != tt.want {
				t.Errorf("addToLastNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addToNextNumber(t *testing.T) {
	type args struct {
		aString string
		x       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{",1],2],3],4]", 8}, ",9],2],3],4]"},
		{"2", args{"]],1]", 2}, "]],3]"},
		{"3", args{"]],0]", 2}, "]],2]"},
		{"4", args{"]],10]", 2}, "]],12]"},
		{"5", args{"]],10", 2}, "]],12"},
		{"6", args{"],0],", 2}, "],2],"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToNextNumber(tt.args.aString, tt.args.x); got != tt.want {
				t.Errorf("addToNextNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_explode(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"[[[[[9,8],1],2],3],4]"}, "[[[[0,9],2],3],4]"},
		{"2", args{"[7,[6,[5,[4,[3,2]]]]]"}, "[7,[6,[5,[7,0]]]]"},
		{"3", args{"[[6,[5,[4,[3,2]]]],1]"}, "[[6,[5,[7,0]]],3]"},
		{"4", args{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"}, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"5", args{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"}, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"6", args{"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"}, "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
		{"7", args{"[[[[0,7],4],[15,[0,13]]],[1,1]]"}, "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := explode(tt.args.aString); got != tt.want {
				t.Errorf("explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digitSplit(t *testing.T) {
	type args struct {
		anumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{10}, "[5,5]"},
		{"2", args{11}, "[5,6]"},
		{"3", args{12}, "[6,6]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitSplit(tt.args.anumber); got != tt.want {
				t.Errorf("digitSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_add(t *testing.T) {
	type fields struct {
		data string
	}
	type args struct {
		b Pair
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Pair
	}{
		{"1", fields{data: "[[[[4,3],4],4],[7,[[8,4],9]]]"}, args{Pair{"[1,1]"}}, Pair{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"}},
		{"2", fields{data: "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"}, args{Pair{"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"}}, Pair{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pair{
				data: tt.fields.data,
			}
			if got := p.add(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pair.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want Pair
	}{

		{"1", args{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]`}, Pair{"[[[[3,0],[5,3]],[4,4]],[5,5]]"}},

		{"2", args{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]
		[6,6]`}, Pair{"[[[[5,0],[7,4]],[5,5]],[6,6]]"}},

		{"3", args{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]
		[6,6]`}, Pair{"[[[[5,0],[7,4]],[5,5]],[6,6]]"}},

		{"4", args{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
		[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
		[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
		[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
		[7,[5,[[3,8],[1,4]]]]
		[[2,[2,2]],[8,[8,1]]]
		[2,9]
		[1,[[[9,3],9],[[9,0],[0,7]]]]
		[[[5,[7,4]],7],1]
		[[[[4,2],2],6],[8,7]]`}, Pair{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}},

		{"5", args{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
		[[[5,[2,8]],4],[5,[[9,9],0]]]
		[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
		[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
		[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
		[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
		[[[[5,4],[7,7]],8],[[8,3],8]]
		[[9,3],[[9,9],[6,[4,9]]]]
		[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
		[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`}, Pair{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.aString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// exploding pair tests
// {"1", args{"[[[[[9,8],1],2],3],4]"}, 9, 8},
// {"2", args{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"}, 7, 3},
// {"3", args{"[7,[6,[5,[4,[3,2]]]]]"}, 3, 2},
// {"4", args{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"}, 3, 2},
// {"5", args{"[[[[0,7],4],[15,[0,13]]],[1,1]]"}, 0, 0},

func Test_getExplodingPair(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{"1", args{"[[[[[9,8],1],2],3],4]"}, "_[9,8", "[[[[_[9,8],1],2],3],4]"},
		{"2", args{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"}, "_[7,3", "[[3,[2,[1,_[7,3]]]],[6,[5,[4,[3,2]]]]]"},
		{"3", args{"[7,[6,[5,[4,[3,2]]]]]"}, "_[3,2", "[7,[6,[5,[4,_[3,2]]]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getExplodingPair(tt.args.aString)
			if got != tt.want {
				t.Errorf("getExplodingPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getExplodingPair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPair_getItems(t *testing.T) {
	type fields struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"1", fields{data: "[[1,2],3]"}, []string{"[1,2]", "3"}},
		{"2", fields{data: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"}, []string{"[[[1,3],[5,3]],[[1,3],[8,7]]]", "[[[4,9],[6,9]],[[8,2],[7,3]]]"}},
		{"3", fields{data: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"}, []string{"[[[6,6],[7,6]],[[7,7],[7,0]]]", "[[[7,7],[7,7]],[[7,8],[9,9]]]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pair{
				data: tt.fields.data,
			}
			if got := p.getItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pair.getItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_magnitude(t *testing.T) {
	type fields struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"-2", fields{"[[3,4],5]"}, 61},
		{"-1", fields{"[9,1]"}, 29},
		{"-2", fields{"[1,9]"}, 21},
		{"0", fields{"[[9,1],[1,9]]"}, 129},
		{"1", fields{"[[1,2],[[3,4],5]]"}, 143},
		{"2", fields{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"}, 1384},
		{"3", fields{"[[[[1,1],[2,2]],[3,3]],[4,4]]"}, 445},
		{"6", fields{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}, 3488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pair{
				data: tt.fields.data,
			}
			if got := p.magnitude(); got != tt.want {
				t.Errorf("Pair.magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFinalSum(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
		[[[5,[2,8]],4],[5,[[9,9],0]]]
		[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
		[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
		[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
		[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
		[[[[5,4],[7,7]],8],[[8,3],8]]
		[[9,3],[[9,9],[6,[4,9]]]]
		[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
		[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`}, 4140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalSum(tt.args.aString); got != tt.want {
				t.Errorf("getFinalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstNumber(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"]],10"}, "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstNumber(tt.args.aString); got != tt.want {
				t.Errorf("getFirstNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundRobinPairs(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
		[[[5,[2,8]],4],[5,[[9,9],0]]]
		[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
		[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
		[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
		[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
		[[[[5,4],[7,7]],8],[[8,3],8]]
		[[9,3],[[9,9],[6,[4,9]]]]
		[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
		[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`}, 3993},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundRobinPairs(tt.args.aString); got != tt.want {
				t.Errorf("roundRobinPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
