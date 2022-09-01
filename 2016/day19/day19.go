package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

type elf struct {
	id   int
	loot int
}

func main() {
	fmt.Println("1.", winningElf(3018458))
	fmt.Println("2.", solvePartTwo(3018458))
}

func winningElf(n int) int {
	elves := []*elf{}
	counter := 0
	for len(elves) < n {
		counter += 1
		elves = append(elves, &elf{counter, 1})
	}
	for len(elves) > 1 {

		for ky, val := range elves {
			if val.loot > 0 {
				if val == elves[len(elves)-1] {
					val.loot += elves[0].loot
					elves[0].loot = 0
					continue
				}
				val.loot += elves[ky+1].loot
				elves[ky+1].loot = 0
			}
		}
		newList := []*elf{}
		for _, val := range elves {
			if val.loot > 0 {
				newList = append(newList, val)
			}
		}
		elves = newList
	}
	return elves[0].id
}

func clockmath(nPositions, position, addend int) int {
	// position += addend
	// for position >= nPositions {
	// 	position -= nPositions
	// }

	return (addend + position) % nPositions
}

func winningElf2(n int) int {
	elves := []*elf{}
	counter := 0
	for len(elves) < n {
		counter += 1
		elves = append(elves, &elf{counter, 1})
	}
	choosingElf := 0
	for len(elves) > 1 {
		elves, choosingElf = removeOne(elves, choosingElf)
	}
	return elves[0].id
}

func winningElf3(n int) int {
	// k is current order; v is original order
	elves := []int{}
	counter := 0
	for len(elves) < n {
		counter += 1
		elves = append(elves, counter)
	}
	currentChooser := 0
	for len(elves) > 1 {
		currentChooser, elves = removeOne2(currentChooser, elves)
	}
	return elves[0] + 1
}

// IntSliceDelete function
func IntSliceDelete(s []int, i int) []int {
	// This creates a new slice by creating 2 slices from the original:
	// s[:i] -> [1, 2]
	// s[i+1:] -> [4, 5, 6]
	// and joining them together using `append`
	return append(s[:i], s[i+1:]...)
}

func removeOne2(currentChooser int, elves []int) (int, []int) {
	length := len(elves)
	if length%100000 == 0 {
		fmt.Println(length)
	}

	halfway := clockmath(length, currentChooser, length/2)

	newItem := IntSliceDelete(elves, halfway)
	return clockmath(length, currentChooser, 1), newItem
}

// func removeOne3(currentChooser int, elves []int) (int, []int) {
// 	length := len(elves)
// 	halfway := clockmath(len(elves), currentChooser, length/2-1)
// 	copy(elves[halfway:], elves[halfway+1])
// 	return clockmath(length-1, currentChooser, 1), newItem
// }

// 26735 too low for second part
// 2136202 too high for second part
// 1509229 also wrong
// correct answer was 1424135

func removeOne(elves []*elf, chooserIndex int) ([]*elf, int) {
	count := len(elves)
	fmt.Println(count)
	halfway := count/2 + 1
	toRemove := clockmath(count, chooserIndex, halfway)
	elves[chooserIndex].loot += elves[toRemove].loot
	newElves := []*elf{}
	for idx, v := range elves {
		if idx != toRemove {
			newElves = append(newElves, v)
		}
	}
	return newElves, chooserIndex
}

// @aceshades on https://www.reddit.com/r/adventofcode/comments/5j4lp1/2016_day_19_solutions/
func solvePartTwo(elfCount int) int {
	// py: pop from back  append to front
	left := deque.Deque[int]{}
	right := deque.Deque[int]{}
	i := 0
	for (left.Len() + right.Len()) < elfCount {
		i += 1
		if i < (elfCount/2)+1 {
			left.PushFront(i)
			continue
		}
		right.PushBack(i)
	}
	for left.Len() > 0 && right.Len() > 0 {
		if left.Len() > right.Len() {
			left.PopBack()
		} else {
			right.PopBack()
		}
		// rotate

		right.PushBack(left.PopBack())
		left.PushFront(right.PopFront())
	}
	fmt.Println(left, right)
	if left.Len() == 1 {
		return left.PopBack()
	}
	return right.PopBack()

}
