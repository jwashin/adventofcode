package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1:", part1(string(input)))
}

type Node struct {
	value int
	idx   int
	next  *Node
	prev  *Node
}

func (a *Node) String() string {
	n := a
	t := []string{}
	idx := a.idx
	for {
		t = append(t, fmt.Sprintf("(%v)", n.value))
		n = n.next
		if n.idx == idx {
			break
		}
	}
	return strings.Join(t, ",")

}

func (n *Node) addNext(k int, x int) *Node {
	g := Node{idx: k, value: x, next: n.next, prev: n}
	n.next.prev = &g
	n.next = &g
	return &g
}

func (n *Node) findIndex(a int) *Node {
	node := n
	for {
		if node.idx == a {
			return node
		}
		node = node.next
	}
}

func (n *Node) findValue(a int) *Node {
	node := n
	for {
		if node.value == a {
			return node
		}
		node = node.next
	}
}

func (n *Node) moveNode() *Node {
	if n.value == 0 {
		return n
	}
	destination := n.toNode(n.value)
	destination.addNext(n.idx, n.value)
	n.delete()
	fmt.Println(destination)
	return destination
}

func (n *Node) delete() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func (n *Node) toNode(a int) *Node {
	out := n
	isNeg := a < 0
	count := abs(a)
	if isNeg {
		// because adding happens to the right
		count += 1
	}
	for t := 0; t < count; t++ {
		if isNeg {
			out = out.prev
		} else {
			out = out.next
		}
	}
	return out
}

func makeNodes(data []int) *Node {

	node := &Node{value: data[0], idx: 0}
	node.next = node
	node.prev = node

	for k, v := range data {
		if k == 0 {
			continue
		}
		node = node.addNext(k, v)
	}
	return node
}

func part1(s string) int {
	input := strings.Split(s, "\n")
	data := []int{}
	for _, v := range input {
		j, _ := strconv.Atoi(strings.TrimSpace(v))
		data = append(data, j)
	}
	node := makeNodes(data)
	fmt.Println(node)
	for i := range data {
		node = node.findIndex(i)
		node = node.moveNode()
	}
	t := node.findValue(0)
	first := t.toNode(1000).value
	second := t.toNode(2000).value
	third := t.toNode(3000).value
	return first + second + third

}
