package main

import "testing"

func TestALU_validate(t *testing.T) {
	type fields struct {
		input          []int
		currInputIndex int
		register       map[string]int
		instructions   []string
	}
	type args struct {
		input        int
		instructions string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ALU{
				input:          tt.fields.input,
				currInputIndex: tt.fields.currInputIndex,
				register:       tt.fields.register,
				instructions:   tt.fields.instructions,
			}
			if got := a.validate(tt.args.input, tt.args.instructions); got != tt.want {
				t.Errorf("ALU.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		a            int
		instructions string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {"1", args{1, `inp w
		// mul x 0
		// add x z
		// mod x 26
		// div z 1
		// add x 11
		// eql x w
		// eql x 0
		// mul y 0
		// add y 25
		// mul y x
		// add y 1
		// mul z y
		// mul y 0
		// add y w
		// add y 7
		// mul y x
		// add z y`}, true},
		{"2", args{1, `inp w
		mul x 0
		add x z
		mod x 26
		add x 14
		eql x w
		eql x 0
		mul y 0
		add y 25
		mul y x
		add y 1
		mul z y
		mul y 0
		add y w
		add y 8
		mul y x
		add z y`}, true},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.a, tt.args.instructions); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
