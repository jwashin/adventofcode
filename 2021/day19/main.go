package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
	z int
}

type Scanner struct {
	id            int
	beaconVectors []Vector
	location      Vector
	orientation   int
}

func parseData(aString string) []Scanner {
	input := strings.Split(aString, "\n")
	var currentScanner Scanner
	scanners := []Scanner{}
	for _, v := range input {
		u := strings.TrimSpace(v)
		if len(u) == 0 {
			scanners = append(scanners, currentScanner)
		} else if strings.Contains(u, "---") {
			var scannerId int
			fmt.Sscanf(u, "--- scanner %d ---", &scannerId)
			currentScanner = Scanner{scannerId, []Vector{}, Vector{0, 0, 0}, 0}
		} else {
			data := strings.Split(u, ",")
			ints := []int{}
			for _, v := range data {
				d, _ := strconv.Atoi(v)
				ints = append(ints, d)
			}
			if len(ints) < 3 {
				ints = append(ints, 0)
			}
			currentScanner.beaconVectors = append(currentScanner.beaconVectors, Vector{ints[0], ints[1], ints[2]})
		}
	}
	if scanners[len(scanners)-1].id != currentScanner.id {
		scanners = append(scanners, currentScanner)
	}
	return scanners
}

func findBeacons(aString string, overlaps int) []Vector {
	scanners := parseData(aString)
	scanner0 := scanners[0]
	for _, v := range scanners[1:] {
		

	}

}
