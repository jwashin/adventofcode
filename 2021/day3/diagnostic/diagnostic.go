package diagnostic

import (
	"fmt"
	"strconv"
)

func MostCommon(aList string) string {

	ones := 0
	zeroes := 1
	for _, val := range aList {
		if val == '1' {
			ones += 1
		}
		if val == '0' {
			zeroes += 1
		}
	}
	if ones > zeroes {
		return "1"
	}
	return "0"
}

func CountOnesAndZeroes(aList string) (int, int) {
	ones := 0
	zeroes := 0
	for _, val := range aList {
		if val == '1' {
			ones += 1
		}
		if val == '0' {
			zeroes += 1
		}
	}
	return ones, zeroes
}

func GetRank(aList []string, rank int) string {
	newList := ""
	for _, val := range aList {
		newList = newList + string(val[rank-1])
	}
	return newList
}

func GetPowerConsumption(aList []string) int {
	gamma := ""
	epsilon := ""
	leastCommon := ""
	mostCommon := ""

	for i := 1; i <= len(aList[0]); i++ {
		mostCommon = MostCommon(GetRank(aList, i))
		leastCommon = "1"
		if mostCommon == "1" {
			leastCommon = "0"
		}

		gamma = gamma + mostCommon
		epsilon = epsilon + leastCommon
	}
	fmt.Println("gamma is " + gamma)
	fmt.Println("epsilon is " + epsilon)

	dGamma, _ := strconv.ParseInt(gamma, 2, 64)
	dEpsilon, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(dGamma) * int(dEpsilon)
}

func NewListFrom(aList []string, position int, value string) []string {
	newList := []string{}
	for _, val := range aList {
		if string(val[position-1]) == value {
			newList = append(newList, val)
		}
	}

	return newList
}

func ReduceMostCommon(aList []string) string {

	for i := 1; i <= len(aList[0]); i++ {
		ones, zeroes := CountOnesAndZeroes(GetRank(aList, i))
		mostCommon := "0"
		if ones >= zeroes {
			mostCommon = "1"
		}
		aList = NewListFrom(aList, i, mostCommon)
		if len(aList) == 1 {
			return aList[0]
		}
	}
	return ""
}

func ReduceLeastCommon(aList []string) string {

	for i := 1; i <= len(aList[0]); i++ {
		ones, zeroes := CountOnesAndZeroes(GetRank(aList, i))
		leastCommon := "0"
		if ones < zeroes {
			leastCommon = "1"
		}
		aList = NewListFrom(aList, i, leastCommon)
		if len(aList) == 1 {
			return aList[0]
		}

	}
	return ""
}

func GetLifeSupportRating(aList []string) int {
	o2 := ReduceMostCommon(aList)
	co2 := ReduceLeastCommon(aList)

	do2, _ := strconv.ParseInt(o2, 2, 64)
	dco2, _ := strconv.ParseInt(co2, 2, 64)
	return int(do2 * dco2)

}
