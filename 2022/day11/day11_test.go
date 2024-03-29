package main

import "testing"

const w = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func Test_monkeyBusiness(t *testing.T) {
	type args struct {
		s      string
		rounds int
		part2  bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{w, 20, false}, 10605},
		{"2", args{w, 10000, true}, 2713310158},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monkeyBusiness(tt.args.s, tt.args.rounds, tt.args.part2); got != tt.want {
				t.Errorf("monkeyBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
