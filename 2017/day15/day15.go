package main

import "fmt"

// puzzle input
var inputA = 289
var inputB = 629

func main() {
	fmt.Println("1.", part1(false))
	fmt.Println("2.", part2(false))
}

type generator struct {
	factor       int
	currentValue int
	xfactor      int
}

func (g *generator) generate() int {
	x := g.currentValue * g.factor % 2147483647
	g.currentValue = x
	return x
}

func (g *generator) next() int {
	d := g.generate()
	for d%g.xfactor != 0 {
		d = g.generate()
	}
	return d
}

func part2(test bool) int {
	startA := inputA
	startB := inputB
	if test {
		startA = 65
		startB = 8921
	}

	a := generator{factor: 16807, currentValue: startA, xfactor: 4}
	ap := &a
	b := generator{factor: 48271, currentValue: startB, xfactor: 8}
	bp := &b
	count := 0
	for trial := 0; trial < 5000000; trial++ {
		t1 := ap.next()
		t2 := bp.next()
		x := fmt.Sprintf("%32b", t1)
		y := fmt.Sprintf("%32b", t2)
		a1 := x[16:]
		a2 := y[16:]
		if a1 == a2 {
			count += 1
		}
	}
	return count
}

func part1(test bool) int {
	startA := inputA
	startB := inputB
	if test {
		startA = 65
		startB = 8921
	}

	a := generator{factor: 16807, currentValue: startA}
	ap := &a
	b := generator{factor: 48271, currentValue: startB}
	bp := &b
	count := 0
	for trial := 0; trial < 40000000; trial++ {
		t1 := ap.generate()
		t2 := bp.generate()
		x := fmt.Sprintf("%32b", t1)
		y := fmt.Sprintf("%32b", t2)
		a1 := x[16:]
		a2 := y[16:]
		if a1 == a2 {
			count += 1
		}
	}
	return count
}
