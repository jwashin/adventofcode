package sonar

import "strconv"

func CountIncreases(aList []string) int {
	previous, _ := strconv.Atoi(aList[0])
	count := 0
	for _, val := range aList {
		newVal, _ := strconv.Atoi(val)
		if newVal > previous {
			count += 1
		}
		previous = newVal
	}
	return count
}

func CountIncreasesbyThrees(aList []string) int {
	intList := []int{}
	for _, val := range aList {
		nv, _ := strconv.Atoi(val)
		intList = append(intList, nv)
	}
	previous := (intList[0] + intList[1] + intList[2])
	count := 0
	lastIdx := len(intList) - 1

	for idx, val := range intList {

		newVal := val
		if (idx + 1) <= lastIdx {
			newVal += intList[idx+1]

		}
		if (idx + 2) <= lastIdx {
			newVal += intList[idx+2]
		} else {
			return count
		}

		if newVal > previous {
			count += 1
		}
		previous = newVal
	}
	return count
}
