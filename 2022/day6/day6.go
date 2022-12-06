package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := string(f)
	fmt.Println("Part 1:", startOfPacket(input, 4))
	fmt.Println("Part 2:", startOfPacket(input, 14))
}

func startOfPacket(s string, n int) int {
	for k := range s {
		item := s[k : k+n]
		if len(item) >= n {
			if allDifferent(item) {
				return k + n
			}
		}
	}
	return 0
}

func allDifferent(s string) bool {
	for k, v := range s {
		for m, h := range s {
			if k != m && v == h {
				return false
			}
		}
	}
	return true
}
