package main

import (
	"testing"
)

func Test_firstParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantName     string
		wantSectorID int
		wantChecksum string
	}{
		{"1", args{"aaaaa-bbb-z-y-x-123[abxyz]"}, "aaaaa-bbb-z-y-x", 123, "abxyz"},
		{"2", args{"a-b-c-d-e-f-g-h-987[abcde]"}, "a-b-c-d-e-f-g-h", 987, "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotSectorID, gotChecksum := firstParse(tt.args.s)
			if gotName != tt.wantName {
				t.Errorf("firstParse() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotSectorID != tt.wantSectorID {
				t.Errorf("firstParse() gotSectorID = %v, want %v", gotSectorID, tt.wantSectorID)
			}
			if gotChecksum != tt.wantChecksum {
				t.Errorf("firstParse() gotChecksum = %v, want %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}

func Test_isRealRoom(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
		want2 int
	}{
		{"1", args{"aaaaa-bbb-z-y-x-123[abxyz]"}, "aaaaa-bbb-z-y-x", true, 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := isRealRoom(tt.args.s)
			if got != tt.want {
				t.Errorf("isRealRoom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isRealRoom() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("isRealRoom() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_shift(t *testing.T) {
	type args struct {
		name string
		id   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"qzmt-zixmtkozy-ivhz", 343}, "very encrypted name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shift(tt.args.name, tt.args.id); got != tt.want {
				t.Errorf("shift() = %v, want %v", got, tt.want)
			}
		})
	}
}
