package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	d, _ := ioutil.ReadFile("input.txt")
	data := strings.TrimSpace(string(d))
	fmt.Println(getVersionSum(data))
	fmt.Println(evaluate(data))
}

type Packet struct {
	version    int
	typeId     int
	value      int
	subPackets []Packet
}

func (p Packet) versionSum() int {
	sum := p.version
	for _, packet := range p.subPackets {
		sum += packet.versionSum()
	}
	return sum
}

func (p *Packet) evaluate() int {
	if p.typeId == 4 {
		// literal
		return p.value
	}
	if p.typeId == 0 {
		// sum
		d := 0
		for _, k := range p.subPackets {
			d += k.evaluate()
		}
		return d
	}
	if p.typeId == 1 {
		// multiply
		d := 1
		for _, k := range p.subPackets {
			d *= k.evaluate()
		}
		return d
	}
	if p.typeId == 2 {
		// max
		d := math.MaxInt
		for _, k := range p.subPackets {
			t := k.evaluate()
			if t < d {
				d = t
			}
		}
		return d
	}
	if p.typeId == 3 {
		// min
		d := math.MinInt
		for _, k := range p.subPackets {
			t := k.evaluate()
			if t > d {
				d = t
			}
		}
		return d
	}
	if p.typeId == 5 {
		// gt
		test1 := p.subPackets[0].evaluate()
		test2 := p.subPackets[1].evaluate()

		if test1 > test2 {
			return 1
		}
		return 0
	}

	if p.typeId == 6 {
		// lt
		test1 := p.subPackets[0].evaluate()
		test2 := p.subPackets[1].evaluate()
		if test1 < test2 {
			return 1
		}
		return 0
	}
	if p.typeId == 7 {
		// eq
		test1 := p.subPackets[0].evaluate()
		test2 := p.subPackets[1].evaluate()
		if test1 == test2 {
			return 1
		}
		return 0
	}

	return 0
}

func (p *Packet) parseData(data string) string {
	if p.typeId == 4 {
		value := ""
		for {
			v := data[0:5]
			data = strings.Replace(data, string(v), "", 1)
			if string(v[0]) == "1" {
				value += v[1:]
			}
			if string(v[0]) == "0" {
				value += v[1:]
				break
			}
		}
		x, _ := strconv.ParseInt(value, 2, 64)
		p.value = int(x)
		return data

	} else {
		// packet is an operator packet
		lengthTypeId := string(data[0])
		if lengthTypeId == "0" {
			// next 15 bits gives total length of subpacket data
			bitLength, _ := strconv.ParseInt(data[1:16], 2, 64)
			theString := string(data[16 : 16+bitLength])
			nextString := data[16+bitLength:]
			for len(theString) > 0 {
				newPacket, s := makePacket(theString)
				theString = s
				p.subPackets = append(p.subPackets, newPacket)
			}
			return nextString

		} else if lengthTypeId == "1" {
			// next 11 bits gives the number of sub-packets
			subPacketCount, _ := strconv.ParseInt(data[1:12], 2, 64)
			theString := string(data[12:])
			for count := 0; count < int(subPacketCount); count++ {
				newPacket, s := makePacket(theString)
				p.subPackets = append(p.subPackets, newPacket)
				theString = s
			}
			return theString

		}

	}
	return data
}

func makePacket(aString string) (Packet, string) {
	b2h := b2HData()
	version, _ := strconv.Atoi(b2h["0"+aString[0:3]])
	ptype, _ := strconv.Atoi(b2h["0"+aString[3:6]])
	packet := Packet{version: version, typeId: ptype}
	s := packet.parseData(aString[6:])
	return packet, s
}

func h2bData() map[string]string {
	h2b := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	return h2b
}
func b2HData() map[string]string {
	h2b := h2bData()
	b2h := map[string]string{}
	for key, val := range h2b {
		b2h[val] = key
	}
	return b2h
}

func getVersionSum(aString string) int {
	packet := decodePacket(aString)
	return packet.versionSum()
}

func evaluate(aString string) int {
	packet := decodePacket(aString)
	return packet.evaluate()
}

func decodePacket(aString string) Packet {
	packet := hexbin(aString)
	z, s := makePacket(packet)
	fmt.Println("Leftovers: " + s)
	return z
}

func hexbin(aString string) string {
	h2b := h2bData()
	s := ""
	for _, val := range aString {
		s = s + h2b[string(val)]
	}
	return s
}
