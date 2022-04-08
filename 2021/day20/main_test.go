package main

import (
	"reflect"
	"testing"
)

func Test_getNine(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"5,10"}, []string{"4,9", "4,10", "4,11", "5,9", "5,10", "5,11", "6,9", "6,10", "6,11"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNine(tt.args.aString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_enhance(t *testing.T) {
	type args struct {
		aString string
		count   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#
		
		#..#.
		#....
		##..#
		..#..
		..###`, 1}, 24},
		{"2", args{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		#..#.
		#....
		##..#
		..#..
		..###`, 2}, 35},
		{"3", args{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		#..#.
		#....
		##..#
		..#..
		..###`, 50}, 3351},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enhance(tt.args.aString, tt.args.count); got != tt.want {
				t.Errorf("enhance() = %v, want %v", got, tt.want)
			}
		})
	}
}
