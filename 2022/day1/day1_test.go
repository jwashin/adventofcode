package main

import (
	"testing"
)

func Test_mostCalories(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`1000
		2000
		3000
		
		4000
		
		5000
		6000
		
		7000
		8000
		9000
		
		10000`}, 24000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCalories(tt.args.data); got != tt.want {
				t.Errorf("mostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostCalories3(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`1000
		2000
		3000
		
		4000
		
		5000
		6000
		
		7000
		8000
		9000
		
		10000`}, 45000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCalories3(tt.args.data); got != tt.want {
				t.Errorf("mostCalories3() = %v, want %v", got, tt.want)
			}
		})
	}
}
