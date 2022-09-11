package main

import (
	"testing"
)

func Test_winningReindeer(t *testing.T) {
	type args struct {
		data string
		time int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"1", args{`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, 1000}, "Comet", 1120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := winningReindeer(tt.args.data, tt.args.time)
			if got != tt.want {
				t.Errorf("winningReindeer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("winningReindeer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_winningReindeer2(t *testing.T) {
	type args struct {
		data string
		time int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"1", args{`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, 1000}, "Dancer", 689},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := winningReindeer2(tt.args.data, tt.args.time)
			if got != tt.want {
				t.Errorf("winningReindeer2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("winningReindeer2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
