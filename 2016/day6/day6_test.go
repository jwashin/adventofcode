package main

import (
	"testing"
)

func Test_maxRepeatsByColumn(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`}, "easter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeatsByColumn(tt.args.a); got != tt.want {
				t.Errorf("maxRepeatsByColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leastRepeatsByColumn(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`}, "advent"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastRepeatsByColumn(tt.args.a); got != tt.want {
				t.Errorf("leastRepeatsByColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
