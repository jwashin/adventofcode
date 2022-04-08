package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(LeastFuelSpent(string(data)))
	fmt.Println(LeastFuelSpent2(string(data)))
}

func LeastFuelSpent(data string) int {
	dataList := strings.Split(strings.TrimSpace(data), ",")
	positions := []int{}
	for _, val := range dataList {
		pos, _ := strconv.Atoi(val)
		positions = append(positions, pos)
	}
	lowestFuelCost := math.MaxInt
	for candidate := min(positions); candidate <= max(positions); candidate++ {
		fuelCost := 0
		for _, val := range positions {
			fc := math.Abs(float64(candidate - val))
			fuelCost += int(fc)
		}
		if fuelCost < lowestFuelCost {
			lowestFuelCost = fuelCost
		}
	}
	return lowestFuelCost

}

func LeastFuelSpent2(data string) int {
	dataList := strings.Split(strings.TrimSpace(data), ",")
	positions := []int{}
	for _, val := range dataList {
		pos, _ := strconv.Atoi(val)
		positions = append(positions, pos)
	}
	lowestFuelCost := math.MaxInt
	for candidate := min(positions); candidate <= max(positions); candidate++ {
		fuelCost := 0
		for _, val := range positions {
			fc := int(math.Abs(float64(candidate - val)))
			fc = triangularSum(fc)
			fuelCost += fc
		}
		if fuelCost < lowestFuelCost {
			lowestFuelCost = fuelCost
		}
	}
	return lowestFuelCost

}

func min(aList []int) int {
	min := aList[0]
	for _, val := range aList {
		if val < min {
			min = val
		}
	}
	return min
}

func max(aList []int) int {
	max := aList[0]
	for _, val := range aList {
		if val > max {
			max = val
		}
	}
	return max
}

func triangularSum(v int) int {
	sum := 0
	for k := 1; k <= v; k++ {
		sum += k
	}
	return sum
}
