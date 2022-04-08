package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := GetInputData("input.txt")
	fmt.Println(CountEasyDigits(data))
	fmt.Println(CalculateOutputTotals(data))
}

func GetInputData(aString string) string {
	dataFile, _ := ioutil.ReadFile(aString)

	return string(dataFile)
}

func CountEasyDigits(aString string) int {
	data := strings.Split(string(aString), "\n")
	count := 0
	for _, line := range data {
		if len(line) > 3 {
			split1 := strings.Split(line, "|")
			s := strings.Split(strings.TrimSpace(split1[1]), " ")

			for _, value := range s {
				length := len(value)
				if length == 2 || length == 4 || length == 3 || length == 7 {
					count += 1
				}
			}
		}
	}
	return count
}

func CalculateOutputTotals(bigString string) int {
	// digits := map[string]string{"abcefg": "0", "cf": "1", "acdeg": "2", "acdfg": "3", "bcdf": "4",
	// 	"adfg": "5", "abdefg": "6", "acf": "7", "abcdefg": "8", "abcdfg": "9"}
	data := strings.Split(string(bigString), "\n")

	count := 0
	for _, line := range data {
		count += lineValue(line)
	}
	return count
}

func lineValue(line string) int {
	if len(line) > 3 {
		split1 := strings.Split(line, "|")
		valueDigits := strings.Split(strings.TrimSpace(split1[1]), " ")

		tenDigits := strings.Split(strings.TrimSpace(split1[0]), " ")
		code := makeCode(tenDigits)
		value := ""
		for _, val := range valueDigits {
			value += decode(val, code)
		}
		v, _ := strconv.Atoi(value)
		return v
	}
	return 0
}

func decode(aString string, decoder map[string]string) string {
	digits := map[string]string{"abcefg": "0", "cf": "1", "acdeg": "2", "acdfg": "3", "bcdf": "4",
		"abdfg": "5", "abdefg": "6", "acf": "7", "abcdefg": "8", "abcdfg": "9"}
	newString := ""
	for _, val := range aString {
		newString += decoder[string(val)]
	}
	x := SortString(newString)
	return digits[x]
}

func makeCode(oldVal []string) map[string]string {
	// oldVal is a list of strings containing the segments of
	// all ten (seven-segment) digits, in an arbitrary order

	f, Two := getSegmentF(oldVal)
	c, One := getSegmentC(oldVal, f)
	a, Seven := getSegmentA(oldVal, c, f)
	Eight, Six, ZeroAndNine := getSegmentBlah(oldVal, c)
	b := getSegmentB(Two, Eight, f)
	d, Four := getSegmentD(oldVal, b, c, f)
	Zero, Nine := GetZeroAndNine(oldVal, ZeroAndNine, d)
	e := GetE(Eight, Nine)
	g := GetG(Two, a, c, d, e)
	fmt.Println(One, Seven, Eight, Six, Four, Six, Nine, Zero)
	return map[string]string{a: "a", b: "b", c: "c", d: "d", e: "e", f: "f", g: "g"}
}

func GetG(Eight string, a string, c string, d string, e string) string {
	temp := Eight
	temp = strings.ReplaceAll(temp, a, "")
	temp = strings.ReplaceAll(temp, c, "")
	temp = strings.ReplaceAll(temp, d, "")
	temp = strings.ReplaceAll(temp, e, "")

	return temp
}

func GetE(Eight string, Nine string) string {
	EightTest := Eight
	for _, val := range Nine {
		EightTest = strings.ReplaceAll(EightTest, string(val), "")
	}
	return EightTest
}

func GetZeroAndNine(aList []string, zeroAndNine []string, d string) (string, string) {
	Zero := ""
	Nine := ""
	for _, val := range zeroAndNine {
		if strings.Contains(val, d) {
			Nine = val
		} else {
			Zero = val
		}
	}
	return Zero, Nine
}

func getSegmentD(aList []string, b string, c string, f string) (string, string) {
	Four := ""
	for _, item := range aList {
		if len(item) == 4 {
			Four = item
			break
		}
	}
	x := Four
	for _, val := range []string{b, c, f} {
		x = strings.ReplaceAll(x, val, "")
	}
	return x, Four

}

func getSegmentB(Two string, Eight string, f string) string {
	TwoTest := Two + f
	B := Eight
	for _, val := range TwoTest {
		B = strings.ReplaceAll(B, string(val), "")
	}
	return B
}

func getSegmentBlah(aList []string, c string) (string, string, []string) {
	ZeroAndSixAndNine := []string{}
	Eight := ""
	Six := ""
	ZeroAndNine := []string{}
	for _, item := range aList {
		if len(item) == 6 {
			ZeroAndSixAndNine = append(ZeroAndSixAndNine, item)
		}
		if len(item) == 7 {
			Eight = item
		}
	}
	for _, item := range ZeroAndSixAndNine {
		if strings.Contains(item, c) {
			ZeroAndNine = append(ZeroAndNine, item)
		} else {
			Six = item
		}
	}
	return Eight, Six, ZeroAndNine
}

func getSegmentA(aList []string, c string, f string) (string, string) {
	for _, item := range aList {
		if len(item) == 3 {
			newItem := item
			Seven := item
			s := strings.ReplaceAll(newItem, f, "")
			a := strings.ReplaceAll(s, c, "")
			return a, Seven
		}
	}
	return "", ""
}

func getSegmentF(aList []string) (string, string) {
	Two := ""
	for _, val := range "abcdefg" {
		v := string(val)
		count := 0
		for _, c := range aList {
			if strings.Contains(c, v) {
				count += 1
			}
		}
		if count == 9 {
			for _, a := range aList {
				if !strings.Contains(a, v) {
					Two = a
				}
			}
			return v, Two
		}
	}
	return "", ""
}

func getSegmentC(aList []string, f string) (string, string) {
	for _, item := range aList {
		if len(item) == 2 {
			newItem := item
			One := item
			s := strings.ReplaceAll(newItem, f, "")
			return s, One
		}
	}
	return "", ""
}

func SortItems(aString []string) []string {
	newList := []string{}
	for _, val := range aString {
		itemList := []string{}
		for _, character := range val {
			item := string(character)
			itemList = append(itemList, item)
		}
		sort.Strings(itemList)
		newString := ""
		for _, val := range itemList {
			newString += val
		}
		newList = append(newList, newString)
	}
	return newList
}

func SortString(aString string) string {
	newString := ""
	newList := []string{}
	for _, char := range aString {
		newList = append(newList, string(char))
	}
	sort.Strings(newList)
	for _, val := range newList {
		newString += val
	}
	return newString
}
