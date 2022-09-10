package main

import "testing"

func Test_getShortestRoute(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`}, 605},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShortestRoute(tt.args.s); got != tt.want {
				t.Errorf("getShortestRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
