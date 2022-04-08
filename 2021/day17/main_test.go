package main

import (
	"reflect"
	"testing"
)

func Test_parseTarget(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name  string
		args  args
		want  Coord
		want1 Coord
	}{
		{"Example", args{"target area: x=20..30, y=-10..-5"}, Coord{20, -5}, Coord{30, -10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTarget(tt.args.aString)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTarget() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseTarget() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_highestY(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"E", args{"target area: x=20..30, y=-10..-5"}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestY(tt.args.aString); got != tt.want {
				t.Errorf("highestY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoord_hitsTarget(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		tl       Coord
		lr       Coord
		velocity Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		{"1", fields{x: 0, y: 0}, args{Coord{20, -5}, Coord{30, -10}, Coord{7, 2}}, true, 3},
		{"2", fields{x: 0, y: 0}, args{Coord{20, -5}, Coord{30, -10}, Coord{6, 3}}, true, 6},
		{"3", fields{x: 0, y: 0}, args{Coord{20, -5}, Coord{30, -10}, Coord{9, 0}}, true, 0},
		{"4", fields{x: 0, y: 0}, args{Coord{20, -5}, Coord{30, -10}, Coord{17, -4}}, false, 0},
		{"5", fields{x: 0, y: 0}, args{Coord{20, -5}, Coord{30, -10}, Coord{6, 9}}, true, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Coord{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			got, got1 := c.hitsTarget(tt.args.tl, tt.args.lr, tt.args.velocity)
			if got != tt.want {
				t.Errorf("Coord.hitsTarget() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Coord.hitsTarget() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_countVelocityChoices(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"E", args{"target area: x=20..30, y=-10..-5"}, 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVelocityChoices(tt.args.aString); got != tt.want {
				t.Errorf("countVelocityChoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
