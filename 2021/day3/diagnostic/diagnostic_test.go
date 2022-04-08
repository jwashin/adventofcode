package diagnostic

import (
	"testing"
)

func TestGetPowerConsumption(t *testing.T) {
	type args struct {
		aList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}}, 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowerConsumption(tt.args.aList); got != tt.want {
				t.Errorf("GetPowerConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLifeSupportRating(t *testing.T) {
	type args struct {
		aList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}}, 230},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLifeSupportRating(tt.args.aList); got != tt.want {
				t.Errorf("GetLifeSupportRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
