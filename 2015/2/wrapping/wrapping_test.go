package wrapping

import (
	"reflect"
	"testing"
)

func TestParseDimensions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		{"parse1", args{"1x2x3"}, &Box{1, 2, 3}},
		{"parse2", args{"12x22x34"}, &Box{12, 22, 34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeBox(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDimensions_PaperTotal(t *testing.T) {
	type fields struct {
		length int
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1", fields{2, 3, 4}, 58},
		{"test2", fields{1, 1, 10}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := Box{
				length: tt.fields.length,
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := dm.PaperTotal(); got != tt.want {
				t.Errorf("Box.PaperTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDimensions_RibbonNeeded(t *testing.T) {
	type fields struct {
		length int
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1", fields{2, 3, 4}, 34},
		{"test2", fields{1, 1, 10}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := Box{
				length: tt.fields.length,
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := dm.RibbonNeeded(); got != tt.want {
				t.Errorf("Box.RibbonNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
