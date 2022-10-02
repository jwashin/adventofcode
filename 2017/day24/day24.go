package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var start = "0"

func main() {
	fmt.Println("1.", part1(false))
	fmt.Println("2.", part2(false))
}

type component struct {
	end1 string
	end2 string
}

func (c *component) asString() string {
	return c.end1 + "/" + c.end2
}

func asComponent(s string) component {
	r := strings.ReplaceAll(s, "/", " ")
	j := strings.Fields(r)
	return component{end1: j[0], end2: j[1]}
}

func (c *component) swap() {
	c.end1, c.end2 = c.end2, c.end1
}

func getItems(s string) []string {
	t := strings.Split(strings.TrimSpace(s), "\n")
	return t
}

type bridge string

func (b bridge) possibles(components []component) []string {
	endPort := b.endPort()
	possibles := []string{}
	for _, v := range components {
		if v.end1 == endPort || v.end2 == endPort {
			if !strings.Contains(string(b), v.end1+"/"+v.end2) {
				possibles = append(possibles, v.asString())
			}
		}
	}
	return possibles
}

func (b bridge) endPort() string {
	components := b.components()
	end := start
	for _, v := range components {
		if v.end1 != end {
			v.swap()
		}
		end = v.end2
	}
	return end
}

var testData = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

func part1(test bool) int {
	input := testData
	if !test {
		zinput, _ := os.ReadFile("input.txt")
		input = string(zinput)
	}
	t := makeBridges(input)
	max := 0
	for _, v := range t {
		t := v.strength()
		if t > max {
			max = t
		}
	}
	return max
}

func part2(test bool) int {
	input := testData
	if !test {
		zinput, _ := os.ReadFile("input.txt")
		input = string(zinput)
	}
	t := makeBridges(input)
	sort.Slice(t, func(i, j int) bool {
		return len(t[i].components()) > len(t[j].components())
	})
	newList := []bridge{}
	length := len(t[0].components())
	for _, v := range t {
		if len(v.components()) == length {
			newList = append(newList, v)
		}
	}
	max := 0
	for _, v := range newList {
		if v.strength() > max {
			max = v.strength()
		}
	}
	return max
}

func (b bridge) strength() int {
	sum := 0
	for _, v := range b.components() {
		v1, _ := strconv.Atoi(v.end1)
		v2, _ := strconv.Atoi(v.end2)
		sum += v1
		sum += v2
	}
	return sum
}

func (b bridge) components() []*component {
	f := strings.Split(string(b), "--")
	items := []*component{}
	for _, item := range f {
		c := asComponent(item)
		items = append(items, &c)
	}
	return items
}

func makeBridges(s string) []bridge {
	components := []component{}
	bridges := []bridge{}
	// var currentItem bridge
	for _, v := range getItems(s) {
		components = append(components, asComponent(v))
	}
	q := []bridge{}
	for _, v := range components {
		if v.end1 == "0" || v.end2 == "0" {
			q = append(q, bridge(v.asString()))
		}
	}
	for len(q) > 0 {
		// sort.Slice(q, func(i, j int) bool {
		// 	return q[i].strength() > q[j].strength()
		// })
		currentItem := q[0]
		bridges = append(bridges, currentItem)
		q = q[1:]
		t := currentItem.possibles(components)
		for _, v := range t {
			q = append(q, bridge(string(currentItem)+"--"+v))
		}
	}
	return bridges
}
