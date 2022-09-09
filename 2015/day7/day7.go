package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	doPart1()
	doPart2()
}

func doPart2() {
	data, _ := os.ReadFile("input2.txt")
	c := makeCircuit(string(data))
	t := c.getSignal("a")
	fmt.Println("1. ", t)
}
func doPart1() {
	data, _ := os.ReadFile("input.txt")
	c := makeCircuit(string(data))
	t := c.getSignal("a")
	fmt.Println("1. ", t)
}

type logicGate struct {
	dest     string
	verb     string
	opa      string
	opb      string
	signal   uint16
	hasValue bool
}

func (gate logicGate) aIsNumber() bool {
	for _, v := range gate.opa {
		if !strings.Contains("0123456789", string(v)) {
			return false
		}
	}
	return true
}

func (gate logicGate) bIsNumber() bool {
	for _, v := range gate.opb {
		if !strings.Contains("0123456789", string(v)) {
			return false
		}
	}
	return true
}

func (gate logicGate) hasSignal() bool {
	return gate.hasValue
}

func (gate logicGate) needsA() bool {
	return gate.opa != "" && !gate.aIsNumber()
}
func (gate logicGate) needsB() bool {
	return gate.opb != "" && !gate.bIsNumber()
}

func (gate *logicGate) makeSignal(inputs map[string]uint16) bool {
	if gate.hasSignal() {
		return false
	}
	var opa, opb uint16
	if gate.needsA() {
		opa = inputs[gate.opa]
	}
	if gate.needsB() {
		opb = inputs[gate.opb]
	}
	if gate.aIsNumber() {
		t, _ := strconv.Atoi(gate.opa)
		opa = uint16(t)
	}
	if gate.bIsNumber() {
		t, _ := strconv.Atoi(gate.opb)
		opb = uint16(t)
	}
	if gate.verb == "ASG" {
		gate.signal = opa
		gate.hasValue = true
		return true
	}
	if gate.verb == "NOT" {
		gate.signal = ^opa
		gate.hasValue = true
		return true
	}
	if gate.verb == "RSHIFT" {
		gate.signal = opa >> opb
		gate.hasValue = true
		return true
	}
	if gate.verb == "LSHIFT" {
		gate.signal = opa << opb
		gate.hasValue = true
		return true
	}
	if gate.verb == "AND" {
		gate.signal = opa & opb
		gate.hasValue = true
		return true
	}
	if gate.verb == "OR" {
		gate.signal = opa | opb
		gate.hasValue = true
		return true
	}
	return false
}

func test(s string) uint16 {
	data := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	c := makeCircuit(data)
	return c.getSignal(s)
}

type circuit map[string]*logicGate

func (c circuit) getSignal(wire string) uint16 {
	v := c[wire]
	if !v.hasValue {
		if v.needsA() && v.needsB() {
			valuea := c.getSignal(v.opa)
			valueb := c.getSignal(v.opb)
			v.makeSignal(map[string]uint16{v.opa: valuea, v.opb: valueb})
			return v.signal
		}
		if v.needsA() {
			valuea := c.getSignal(v.opa)
			v.makeSignal(map[string]uint16{v.opa: valuea})
			return v.signal
		}
		if v.needsB() {
			valueb := c.getSignal(v.opb)
			v.makeSignal(map[string]uint16{v.opb: valueb})
			return v.signal
		}
	}
	return v.signal
}

func makeGate(s string) *logicGate {
	t := strings.Split(s, " ")
	dest := t[len(t)-1]
	ops := t[:len(t)-2]

	if len(ops) == 1 {
		value, err := strconv.Atoi(ops[0])
		if err == nil {
			return &logicGate{signal: uint16(value), dest: dest, hasValue: true}
		} else {
			return &logicGate{opa: ops[0], verb: "ASG", dest: dest}
		}
	}
	if len(ops) == 2 {

		return &logicGate{verb: ops[0], opa: ops[1], dest: dest}
	}
	if len(ops) == 3 {
		return &logicGate{verb: ops[1], opa: ops[0], opb: ops[2], dest: dest}
	}
	return &logicGate{dest: dest}
}

func makeCircuit(s string) circuit {
	items := circuit{}
	for _, v := range strings.Split(s, "\n") {
		item := makeGate(v)
		items[item.dest] = item
	}
	return items
}
