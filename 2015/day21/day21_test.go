package main

import (
	"testing"
)

func Test_playGame(t *testing.T) {
	type args struct {
		player player
		boss   player
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{player: player{hitPoints: 8, items: []item{{damage: 5}, {armor: 5}}},
			boss: player{hitPoints: 12, items: []item{{damage: 7}, {armor: 2}}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.player, tt.args.boss); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
		want2 int
		want3 int
	}{
		{"1", args{"Dagger        8     4       0"}, "Dagger", 8, 4, 0},
		{"2", args{"Defense +3   80     0       3"}, "Defense+3", 80, 0, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := storeParse(tt.args.s)
			if got != tt.want {
				t.Errorf("storeParse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("storeParse() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("storeParse() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("storeParse() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
