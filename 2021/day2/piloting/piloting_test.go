package piloting

import "testing"

func TestNavigate(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test case", args{[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}}, 900}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Navigate(tt.args.instructions); got != tt.want {
				t.Errorf("Navigate() = %v, want %v", got, tt.want)
			}
		})
	}
}
