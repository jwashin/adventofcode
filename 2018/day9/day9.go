package main

import "fmt"

// var puzzleInput = "493 players; last marble is worth 71863 points"

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

func (n *Node) previous() *Node {
	if n.prev != nil {
		return n.prev
	}
	node := n
	for node.next != n {
		node = node.next
	}
	return node
}

func (n *Node) delete() *Node {
	t := n.previous()
	t.next = n.next
	return n.next
}

func main() {
	fmt.Println("1.", part1a(493, 71863))
	fmt.Println("2.", part1a(493, 7186300))
}

type marbles struct {
	circle        []int
	currentMarble int
}

func (m *marbles) add(n int) int {
	if n%23 == 0 {
		return m.do23(n)
	}
	t := m.currentMarble
	for x := 0; x < 1; x++ {
		t += 1
		if t > len(m.circle)-1 {
			t = 0
		}
	}
	insertionPoint := t + 1
	first := m.circle[:insertionPoint]
	last := m.circle[insertionPoint:]
	// s.insertion += 1
	b2 := []int{}
	b2 = append(b2, first...)
	b2 = append(b2, n)
	b2 = append(b2, last...)
	m.currentMarble = insertionPoint
	m.circle = b2
	return 0
}

func (m *marbles) do23(n int) int {
	score := n
	t := m.currentMarble
	for x := 0; x < 7; x++ {
		t -= 1
		if t < 0 {
			t = len(m.circle) - 1
		}
	}
	insertionPoint := t
	score += m.circle[t]
	first := m.circle[:insertionPoint]
	last := m.circle[insertionPoint+1:]
	b2 := []int{}
	b2 = append(b2, first...)
	b2 = append(b2, last...)
	m.currentMarble = insertionPoint
	m.circle = b2
	return score
}

func part1(players int, hiValue int) int {
	marbles := marbles{circle: []int{0}, currentMarble: 0}
	n := 0
	scores := map[int]int{}

	for n <= hiValue {
		player := n % players
		n += 1
		scores[player] += marbles.add(n)

	}

	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max

}
func part1a(players int, hiValue int) int {
	currentMarble := &Node{value: 0}
	currentMarble.next = currentMarble
	currentMarble.prev = currentMarble
	n := 0
	scores := map[int]int{}
	for n <= hiValue {
		n += 1
		player := n % players
		if n%1000000 == 0 {
			fmt.Println(n, "of", hiValue)
		}
		if n%23 != 0 {
			currentMarble = currentMarble.next.addNext(n)
			continue
		}
		for i := 0; i < 7; i++ {
			currentMarble = currentMarble.previous()
		}
		scores[player] += n + currentMarble.value
		currentMarble = currentMarble.delete()
	}
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max

}
