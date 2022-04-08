package main

import (
	"testing"
)

func Test_countPaths(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{`start-A
		start-b
		A-c
		A-b
		b-d
		A-end
		b-end`}, 10},
		{"Example2", args{`dc-end
		HN-start
		start-kj
		dc-start
		dc-HN
		LN-dc
		HN-end
		kj-sa
		kj-HN
		kj-dc`}, 19},
		{"Example3", args{`fs-end
		he-DX
		fs-he
		start-DX
		pj-DX
		end-zg
		zg-sl
		zg-pj
		pj-he
		RW-he
		fs-DX
		pj-RW
		zg-RW
		start-pj
		he-WI
		zg-he
		pj-fs
		start-RW`}, 226},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPaths(tt.args.aString); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countImprovedPaths(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{`start-A
		start-b
		A-c
		A-b
		b-d
		A-end
		b-end`}, 36},
		{"Example2", args{`dc-end
		HN-start
		start-kj
		dc-start
		dc-HN
		LN-dc
		HN-end
		kj-sa
		kj-HN
		kj-dc`}, 103},
		{"Example3", args{`fs-end
		he-DX
		fs-he
		start-DX
		pj-DX
		end-zg
		zg-sl
		zg-pj
		pj-he
		RW-he
		fs-DX
		pj-RW
		zg-RW
		start-pj
		he-WI
		zg-he
		pj-fs
		start-RW`}, 3509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countImprovedPaths(tt.args.aString); got != tt.want {
				t.Errorf("countImprovedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
