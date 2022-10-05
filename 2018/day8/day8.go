package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part1(false))
	fmt.Println("2.", part2(false))
}

type node struct {
	childCount    int
	metadataCount int
	children      []node
	metadata      []int
}

func part1(test bool) int {
	d := getData(test)
	h, _ := nodeFromList(d)
	return h.sumMetadata()
}

func part2(test bool) int {
	d := getData(test)
	h, _ := nodeFromList(d)
	return h.value()
}

func getData(test bool) []int {
	data := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

	if !test {
		input, _ := os.ReadFile("input.txt")
		data = string(input)
	}
	f := strings.Fields(data)
	d := []int{}
	for _, v := range f {
		t, _ := strconv.Atoi(v)
		d = append(d, t)
	}
	return d
}

func (n node) sumMetadata() int {
	sum := 0
	for _, v := range n.metadata {
		sum += v
	}
	for _, v := range n.children {
		sum += v.sumMetadata()
	}
	return sum
}

func (n node) value() int {
	sum := 0
	if n.childCount == 0 {
		for _, v := range n.metadata {
			sum += v
		}
		return sum
	}
	for _, v := range n.metadata {
		if v > 0 && v <= n.childCount {
			sum += n.children[v-1].value()
		}
	}
	return sum
}

func nodeFromList(s []int) (node, []int) {
	d := node{childCount: s[0], metadataCount: s[1]}
	s = s[2:]
	for i := 0; i < d.childCount; i++ {
		var child node
		child, s = nodeFromList(s)
		d.children = append(d.children, child)
	}
	d.metadata = s[:d.metadataCount]
	s = s[d.metadataCount:]
	return d, s
}
