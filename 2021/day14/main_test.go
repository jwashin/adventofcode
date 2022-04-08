package main

import (
	"testing"
)

func Test_doPairInsert(t *testing.T) {
	type args struct {
		data  string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 1}, "NCNBCHB"},
		{"Example 2", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 2}, "NBCCNBBBCBHCB"},
		{"Example 4", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 4}, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"},
		{"Example 1a", args{`NN

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 2}, "NBCCN"},
		{"Example 1b", args{`NN

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 3}, "NBBBCNCCN"},
		{"Example 1c", args{`NN

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 4}, "NBBNBNBBCCNBCNCCN"},
		{"Example 3a", args{`NC

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 1}, "NBC"},
		{"Example 3b", args{`NC

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 2}, "NBBBC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPairInsert(tt.args.data, tt.args.count); got != tt.want {
				t.Errorf("doPairInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTestValue(t *testing.T) {
	type args struct {
		data  string
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 10", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 10}, 1588},
		// {"Example 10a", args{`NN

		// CH -> B
		// HH -> N
		// CB -> H
		// NH -> C
		// HB -> C
		// HC -> B
		// HN -> C
		// NN -> C
		// BH -> H
		// NC -> B
		// NB -> B
		// BN -> B
		// BB -> N
		// BC -> B
		// CC -> N
		// CN -> C`, 40}, 2188189693529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTestValue(tt.args.data, tt.args.count); got != tt.want {
				t.Errorf("getTestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDistribution(t *testing.T) {
	type args struct {
		data  string
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"Example 10", args{`NNCB

		// CH -> B
		// HH -> N
		// CB -> H
		// NH -> C
		// HB -> C
		// HC -> B
		// HN -> C
		// NN -> C
		// BH -> H
		// NC -> B
		// NB -> B
		// BN -> B
		// BB -> N
		// BC -> B
		// CC -> N
		// CN -> C`, 10}, 1588},
		{"Example 10a", args{`NB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 40}, 2188189693529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeDistribution(tt.args.data, tt.args.count); got != tt.want {
				t.Errorf("makeDistribution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDistributionCounts(t *testing.T) {
	type args struct {
		data  string
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 10", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 10}, 1588},
		{"Example 10a", args{`NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C`, 40}, 2188189693529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeDistributionCounts(tt.args.data, tt.args.count); got != tt.want {
				t.Errorf("makeDistributionCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
