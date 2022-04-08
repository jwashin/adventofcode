package lines

import "testing"

func TestCountIntersections(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIntersections(tt.args.input); got != tt.want {
				t.Errorf("CountIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}
