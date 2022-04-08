package lanternfish

import "testing"

func TestCountFish(t *testing.T) {
	type args struct {
		aString string
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{"3,4,3,1,2", 18}, 26},
		{"Example2", args{"3,4,3,1,2", 80}, 5934},
		{"Example3", args{"3,4,3,1,2", 256}, 26984457539},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFish(tt.args.aString, tt.args.days); got != tt.want {
				t.Errorf("CountFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
