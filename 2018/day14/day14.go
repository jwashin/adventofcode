package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("1.", part1(306281))
	fmt.Println("2.", part2("306281"))
}

// circular linked list
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) addNext(x int) *Node {
	g := Node{value: x, next: n.next, prev: n}
	n.next.prev = &g
	n.next = &g
	return &g
}

func (n *Node) delete() *Node {
	t := n.prev
	t.next = n.next
	return n.next
}

func part2(index string) int {
	elf1 := &Node{value: 3}
	elf1.next = elf1
	elf1.prev = elf1
	elf2 := elf1.addNext(7)
	recipeCount := 2
	end := elf2
	key := index
	test := "37"
	for {
		v := elf1.value + elf2.value
		s := fmt.Sprint(v)
		for _, v := range s {
			v2, _ := strconv.Atoi(string(v))
			end = end.addNext(v2)
			test += fmt.Sprint(v2)
			for len(test) > len(index) {
				test = test[1:]
			}
			recipeCount += 1
			if test == key {
				return recipeCount - len(test)
			}
		}
		elf1Move := elf1.value + 1
		elf2Move := elf2.value + 1
		for m := 0; m < elf1Move; m++ {
			elf1 = elf1.next
		}
		for m := 0; m < elf2Move; m++ {
			elf2 = elf2.next
		}
	}
}

func part1(index int) string {
	elf1 := &Node{value: 3}
	elf1.next = elf1
	elf1.prev = elf1
	elf2 := elf1.addNext(7)
	recipeCount := 2
	end := elf2
	out := ""
	for {
		v := elf1.value + elf2.value
		s := fmt.Sprint(v)
		for _, v := range s {
			v2, _ := strconv.Atoi(string(v))
			end = end.addNext(v2)
			recipeCount += 1
			if recipeCount > index {
				out += string(v)
			}
			if len(out) == 10 {
				return out
			}
		}
		elf1Move := elf1.value + 1
		elf2Move := elf2.value + 1
		for m := 0; m < elf1Move; m++ {
			elf1 = elf1.next
		}
		for m := 0; m < elf2Move; m++ {
			elf2 = elf2.next
		}
	}
}
