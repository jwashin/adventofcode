package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1.", part2())
}

func part1() string {
	input, _ := os.ReadFile("input.txt")
	nodes := getNodes(string(input))
	return findBaseProgram(nodes)
}

func part2() string {
	input, _ := os.ReadFile("input.txt")
	nodes := getNodes(string(input))
	z := findBaseProgram(nodes)
	oddballBranch := nodes[z].oddballBranch(nodes)
	for oddballBranch != "" {
		oddballBranch = nodes[oddballBranch].oddballBranch(nodes)
		fmt.Println("")
	}
	return nodes[oddballBranch].parent

	// iterated through the debugger
	// and the input source to get the answer, 910
}

type node struct {
	name     string
	value    int
	parent   string
	children string
}

func (n *node) oddballBranch(z map[string]*node) string {
	d := map[int]int{}
	t := n.childWeights(z)
	oddballWeight := 0
	for _, v := range t {
		d[v] += 1
	}
	for k, v := range d {
		if v == 1 {
			oddballWeight = k
		}
	}
	for k, v := range t {
		if v == oddballWeight {
			return k
		}
	}
	return ""
}

func (n *node) weight(z map[string]*node) int {
	w := n.value
	if len(n.children) > 0 {
		names := strings.Split(n.children, ",")
		for _, v := range names {
			w += z[v].weight(z)
		}
	}
	fmt.Println(n.name, w)
	return w
}

func (n *node) childWeights(z map[string]*node) map[string]int {
	if len(n.children) == 0 {
		return map[string]int{}
	}
	t := strings.Split(n.children, ",")
	w := map[string]int{}
	for _, v := range t {
		w[v] = z[v].weight(z)
	}
	return w
}

func test(s string) string {
	nodes := getNodes(s)
	return findBaseProgram(nodes)
}

func findBaseProgram(nodes map[string]*node) string {
	for k, v := range nodes {
		if v.parent == "" {
			return k
		}
	}
	return "not found"
}

func getNodes(s string) map[string]*node {
	data := strings.Split(s, "\n")

	nodes := map[string]*node{}

	for _, v := range data {
		item := node{}
		fields := strings.Fields(v)
		item.name = fields[0]
		item.value, _ = strconv.Atoi(fields[1][1 : len(fields[1])-1])
		item.children = ""
		if len(fields) > 2 {

			if fields[2] == "->" {
				children := []string{}
				for _, v := range fields[3:] {
					t := strings.Replace(v, ",", "", 1)
					children = append(children, t)
				}
				item.children = strings.Join(children, ",")
			}
		}
		nodes[item.name] = &item
	}
	for _, v := range nodes {
		if len(v.children) > 0 {
			for _, child := range strings.Split(v.children, ",") {
				c := nodes[child]
				c.parent = v.name
			}
		}
	}
	return nodes
}
