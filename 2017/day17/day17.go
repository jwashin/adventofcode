package main

import "fmt"

var puzzleInput = 316

func main() {
	fmt.Println("1.", part1(false))
	fmt.Println("2.", part2(false))
}

type spinlock struct {
	increment       int
	insertion       int
	currentPosition int
	buffer          []int
}

func (s *spinlock) insert() {
	t := s.currentPosition
	for x := 0; x < s.increment; x++ {
		t += 1
		if t > len(s.buffer)-1 {
			t = 0
		}
	}
	insertionPoint := t + 1
	first := s.buffer[:insertionPoint]
	last := s.buffer[insertionPoint:]
	s.insertion += 1
	b2 := []int{}
	b2 = append(b2, first...)
	b2 = append(b2, s.insertion)
	b2 = append(b2, last...)
	s.currentPosition = insertionPoint
	s.buffer = b2
}
func (s *spinlock) insert2() {
	// if we only care about the value in buffer[1]
	t := s.currentPosition
	for x := 0; x < s.increment; x++ {
		t += 1
		if t > s.insertion {
			t = 0
		}
	}
	insertionPoint := t + 1
	s.insertion += 1

	if t+1 == 1 {
		first := s.buffer[:insertionPoint]
		last := s.buffer[insertionPoint:]
		b2 := []int{}
		b2 = append(b2, first...)
		b2 = append(b2, s.insertion)
		b2 = append(b2, last...)
		s.buffer = b2
	}
	s.currentPosition = insertionPoint
}

// 1785 too high
func part1(test bool) int {
	increment := puzzleInput
	if test {
		increment = 3
	}

	z := spinlock{increment: increment, insertion: 0, currentPosition: 0, buffer: []int{0}}
	for i := 0; i < 2017; i++ {
		z.insert()
	}
	fmt.Println(z.buffer)
	return z.buffer[z.currentPosition+1]
}

// 53228 too low
func part2(test bool) int {
	increment := puzzleInput
	if test {
		increment = 3
	}

	z := spinlock{increment: increment, insertion: 0, currentPosition: 0, buffer: []int{0}}
	bufferValue := 0
	for i := 0; i < 50000000; i++ {
		z.insert2()
		if z.buffer[1] != bufferValue {
			bufferValue = z.buffer[1]
			fmt.Println(i, z.buffer[1])
		}
	}

	return z.buffer[1]
}

func test(incrementCount int) (int, []int) {
	z := spinlock{increment: 3, insertion: 0, currentPosition: 0, buffer: []int{0}}
	for i := 0; i < incrementCount; i++ {
		z.insert()
	}
	return z.currentPosition, z.buffer
}
