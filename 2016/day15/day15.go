package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := os.ReadFile("input.txt")
	fmt.Println(doPush(string(s)))
	fmt.Println(doPush2(string(s)))
}

func clockmath(nPositions, position, addend int) int {
	// position += addend
	// for position >= nPositions {
	// 	position -= nPositions
	// }

	return (addend + position) % nPositions
}

func parseInput(s string) (int, int, int) {
	var discNumber, nPos, pos int
	fmt.Sscanf(s, "Disc #%d has %d positions; at time=0, it is at position %d.", &discNumber, &nPos, &pos)
	return discNumber, nPos, pos
}

type disk struct {
	id         int
	nPositions int
	position   int
}

func (d disk) positionAt(t int) int {
	return clockmath(d.nPositions, d.position, t)
}

func getDiscs(s string) []disk {
	t := strings.Split(s, "\n")
	// discs := []disk{{0, 1, 0}}
	discs := []disk{}

	for _, v := range t {
		d, n, p := parseInput(v)
		discs = append(discs, disk{id: d, nPositions: n, position: p})
	}
	return discs
}

func doPush(s string) int {
	disks := getDiscs(s)
	// positions := []int{}
	for i, j := 0, len(disks)-1; i < j; i, j = i+1, j-1 {
		disks[i], disks[j] = disks[j], disks[i]
	}
	won := false
	t := 0
	for !won {
		t += 1
		if t > len(disks) {
			won = true
			for ky, disc := range disks {
				if disc.positionAt(t-ky) != 0 {
					won = false
					break
				}
			}

		}
	}
	return t - len(disks)
}
func doPush2(s string) int {
	disks := getDiscs(s)
	// positions := []int{}
	disks = append(disks, disk{id: 8, nPositions: 11, position: 0})
	for i, j := 0, len(disks)-1; i < j; i, j = i+1, j-1 {
		disks[i], disks[j] = disks[j], disks[i]
	}
	won := false
	t := 0
	for !won {
		t += 1
		if t > len(disks) {
			won = true
			for ky, disc := range disks {
				if disc.positionAt(t-ky) != 0 {
					won = false
					break
				}
			}

		}
	}
	return t - len(disks)
}
