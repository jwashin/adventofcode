package main

import (
	"testing"
)

// func Test_layer_scannerLocation(t *testing.T) {
// 	type fields struct {
// 		depth int
// 		rge   int
// 	}
// 	type args struct {
// 		time int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   int
// 	}{
// 		{"1", fields{0, 3}, args{0}, 0},
// 		{"2", fields{0, 3}, args{1}, 1},
// 		{"3", fields{0, 3}, args{2}, 2},
// 		{"4", fields{0, 3}, args{3}, 1},
// 		{"5", fields{0, 3}, args{4}, 0},
// 		{"6", fields{0, 4}, args{4}, 2},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			j := layer{
// 				depth: tt.fields.depth,
// 				rge:   tt.fields.rge,
// 			}
// 			if got := j.scannerLocation(tt.args.time); got != tt.want {
// 				t.Errorf("layer.scannerLocation() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_sumSeverity(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0: 3
		1: 2
		4: 4
		6: 4`}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSeverity(tt.args.s); got != tt.want {
				t.Errorf("sumSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcDelay(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0: 3
		1: 2
		4: 4
		6: 4`}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDelay(tt.args.s); got != tt.want {
				t.Errorf("calcDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcDelay2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`0: 3
		1: 2
		4: 4
		6: 4`}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDelay2(tt.args.s); got != tt.want {
				t.Errorf("calcDelay2() = %v, want %v", got, tt.want)
			}
		})
	}
}
