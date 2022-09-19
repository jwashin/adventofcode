package main

import (
	"reflect"
	"testing"
)

// {"1", args{player: player{hitPoints: 10, mana: 250}, boss: boss{hitPoints: 13, damage: 8}}, true, false, timerList{}},

// {"1", args{player: &player{hitPoints: 10, mana: 250}, boss: &boss{hitPoints: 13, damage: 8}, plays: "pm"}, true, false, timerList{"p": 3}},
// {"2", args{player: &player{hitPoints: 10, mana: 250}, boss: &boss{hitPoints: 14, damage: 8}, plays: "rsdpm"}, true, false, timerList{"p": 3}},

func Test_playGame2(t *testing.T) {
	type args struct {
		player *player
		boss   *boss
		plays  string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
		want2 timerList
	}{
		// {"1", args{player: &player{hitPoints: 10, mana: 250}, boss: &boss{hitPoints: 13, damage: 8}, plays: "pm"}, true, false, timerList{"p": 3}},
		// {"2", args{player: &player{hitPoints: 10, mana: 250}, boss: &boss{hitPoints: 14, damage: 8}, plays: "rsdpm"}, true, false, timerList{"p": 3}},
		{"3", args{player: &player{hitPoints: 15, mana: 250}, boss: &boss{hitPoints: 14, damage: 8}, plays: "rsdpm"}, true, false, timerList{"p": 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := playGame2(tt.args.player, tt.args.boss, tt.args.plays)
			if got != tt.want {
				t.Errorf("playGame2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("playGame2() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("playGame2() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
