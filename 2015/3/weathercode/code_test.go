package weathercode

import (
	"reflect"
	"testing"
)

func TestNextValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{20151125}, 31916031},
		{"second", args{31916031}, 18749137},
		{"6,1", args{33071741}, 17552253},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextValue(tt.args.value); got != tt.want {
				t.Errorf("NextValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextLocation(t *testing.T) {
	type args struct {
		old Coordinate
	}
	tests := []struct {
		name string
		args args
		want Coordinate
	}{
		{"first", args{Coordinate{1, 1, 1, 0}}, Coordinate{2, 2, 1, 0}},
		{"second", args{Coordinate{2, 2, 1, 0}}, Coordinate{2, 1, 2, 0}},
		{"third", args{Coordinate{2, 1, 2, 0}}, Coordinate{3, 3, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextLocation(tt.args.old); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
