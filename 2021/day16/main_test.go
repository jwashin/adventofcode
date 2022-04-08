package main

import (
	"reflect"
	"testing"
)

func Test_hexbin(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"D2FE28"}, "110100101111111000101000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexbin(tt.args.aString); got != tt.want {
				t.Errorf("hexbin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodePacket(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want Packet
	}{
		{"Example 1", args{"D2FE28"}, Packet{version: 6, typeId: 4, value: 2021}},
		{"Example 2", args{"38006F45291200"}, Packet{version: 1, typeId: 6, value: 0,
			subPackets: []Packet{{version: 6, typeId: 4, value: 10},
				{version: 2, typeId: 4, value: 20}}}},
		{"Example 3", args{"EE00D40C823060"}, Packet{version: 7, typeId: 3, value: 0,
			subPackets: []Packet{{version: 2, typeId: 4, value: 1},
				{version: 4, typeId: 4, value: 2},
				{version: 1, typeId: 4, value: 3},
			}}},
		{"Example 4", args{"620080001611562C8802118E34"},
			Packet{version: 3, typeId: 0, value: 0,
				subPackets: []Packet{
					{version: 0, typeId: 0, value: 0,
						subPackets: []Packet{
							{version: 0, typeId: 4, value: 10},
							{version: 5, typeId: 4, value: 11},
						},
					},
					{version: 1, typeId: 0, value: 0,
						subPackets: []Packet{
							{version: 0, typeId: 4, value: 12},
							{version: 3, typeId: 4, value: 13},
						},
					},
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodePacket(tt.args.aString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVersionSum(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"8A004A801A8002F478"}, 16},
		{"example 2", args{"620080001611562C8802118E34"}, 12},
		{"example 3", args{"C0015000016115A2E0802F182340"}, 23},
		{"example 4", args{"A0016C880162017C3686B18A3D4780"}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVersionSum(tt.args.aString); got != tt.want {
				t.Errorf("getVersionSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacket_evaluate(t *testing.T) {
	type fields struct {
		version    int
		typeId     int
		value      int
		subPackets []Packet
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Packet{
				version:    tt.fields.version,
				typeId:     tt.fields.typeId,
				value:      tt.fields.value,
				subPackets: tt.fields.subPackets,
			}
			if got := p.evaluate(); got != tt.want {
				t.Errorf("Packet.evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluate(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"C200B40A82"}, 3},
		{"example2", args{"04005AC33890"}, 54},
		{"example3", args{"880086C3E88112"}, 7},
		{"example4", args{"CE00C43D881120"}, 9},
		{"example5", args{"D8005AC2A8F0"}, 1},
		{"example6", args{"F600BC2D8F"}, 0},
		{"example7", args{"9C005AC2F8F0"}, 0},
		{"example8", args{"9C0141080250320F1802104A08"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.aString); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
