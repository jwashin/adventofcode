package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pair struct {
	data string
}

func (p *Pair) add(b Pair) Pair {
	s := fmt.Sprintf("[%s,%s]", p.data, b.data)
	s = reduce(s)
	return Pair{s}
}

func (p *Pair) getItems() []string {
	level := 0
	checkLevel := 1
	pair := []string{}
	for i, v := range p.data {
		c := string(v)
		if c == "[" {
			level += 1
		}
		if c == "]" {
			level -= 1
		}
		if level == checkLevel && c == "," {
			g := p.data[1:i]
			h := p.data[i+1 : len(p.data)-1]
			if len(g) > 0 {
				pair = append(pair, g)
			}
			if len(h) > 0 {
				pair = append(pair, h)
			}
			return pair
		}
	}
	fmt.Println("there was something wrong with " + p.data)
	return pair
}

func (p *Pair) magnitude() int {
	theList := p.getItems()
	if len(theList) == 0 {
		fmt.Println("No values in " + p.data)
		return 0
	}
	if len(theList) == 1 {
		fmt.Println("No x values in " + p.data)
		return 0
	}
	a := theList[0]
	b := theList[1]
	var aVal, bVal int
	if isDigit(a) {
		aVal, _ = strconv.Atoi(a)
	} else {
		newItem := Pair{a}
		aVal = newItem.magnitude()
	}
	if isDigit(b) {
		bVal, _ = strconv.Atoi(b)
	} else {
		newItem := Pair{b}
		bVal = newItem.magnitude()
	}
	return 3*aVal + 2*bVal
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	fmt.Println(getFinalSum(string(data)))
	fmt.Println(roundRobinPairs(string(data)))
}

func roundRobinPairs(aString string) int {
	theList := strings.Split(aString, "\n")
	for _, v := range theList {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			theList = append(theList, v)
		}
	}
	max := 0
	for _, y := range theList {
		for _, x := range theList {
			if x != y {
				v1 := Pair{y}
				s1 := v1.add(Pair{x})
				ts := s1.magnitude()
				if ts > max {
					max = ts
				}
				v2 := Pair{x}
				s2 := v2.add(Pair{y})
				ts1 := s2.magnitude()
				if ts1 > max {
					max = ts1
				}
			}
		}
	}
	return max
}

func getFinalSum(aString string) int {
	it := sum(aString)
	return it.magnitude()
}

func parsePair(aString string) (string, string) {
	// openers := "([{<"
	s := aString[1 : len(aString)-1]
	openers := getLevel(s)
	closers := countClosers(s)

	if openers == closers && closers == 0 {
		a := strings.Split(s, ",")
		return a[0], a[1]
	}

	if string(s[0]) == "[" {
		a := findCloser(s)
		b := s[0 : a+1]
		c := s[a+2:]
		return b, c
	} else {
		a := strings.IndexByte(s, ',')
		b := s[0:a]
		c := s[a+1:]
		return b, c
	}
}

func sum(aString string) Pair {
	theList := strings.Split(aString, "\n")
	item := Pair{strings.TrimSpace(theList[0])}
	for _, v := range theList[1:] {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			item = item.add(Pair{v})
		}
	}
	return item
}

func getLevel(aString string) int {
	level := 0
	maxLevel := 0
	for _, v := range aString {
		c := string(v)
		if c == "[" {
			level += 1
			if level > maxLevel {
				maxLevel = level
			}
		}
		if c == "]" {
			level -= 1
		}
	}
	return maxLevel
}

func findCloser(aString string) int {
	closeCount := 0
	for i, v := range aString {
		s := string(v)
		if s == "[" {
			closeCount += 1
		}
		if s == "]" {
			closeCount -= 1
			if closeCount == 0 {
				return i
			}
		}
	}
	return 0
}

func countClosers(aString string) int {

	return countChar(aString, "]")
}

func countChar(aString string, t string) int {
	count := 0
	for _, v := range aString {
		if string(v) == t {
			count += 1
		}
	}
	return count
}

func reduce(aString string) string {
	d2 := aString
	if getLevel(d2) > 4 {
		d2 := explode(d2)
		return reduce(d2)
	}

	j := digitOver10(d2)
	if j >= 10 {
		d2 = split(d2, j)
		return reduce(d2)
	}
	return d2
}

func split(aString string, anum int) string {
	ds := digitSplit(anum)

	newString := strings.Replace(aString, fmt.Sprint(anum), ds, 1)

	return newString
}

func getExplodingPair(aString string) (string, string) {
	newStr := ""
	level := 0
	rep := ""
	newString := ""

	for idx, v := range aString {
		c := string(v)
		if c == "[" {
			level += 1
			rep = "_["
		}
		if c == "]" {
			if len(newStr) > 0 {
				remainder := newString + aString[idx:]
				return newStr, remainder
			}
			level -= 1
			newStr = ""
		}
		if level == 5 {
			if c == "[" {
				newStr = newStr + rep
				newString += rep
				continue
			}
			newStr = newStr + c
		}
		newString += c

	}

	return newStr, aString

}

func explode(aString string) string {
	var p1, p2 int
	s1, s2 := getExplodingPair(aString)
	fmt.Sscanf(s1, "_[%d,%d", &p1, &p2)
	// pr := fmt.Sprintf("_[%d,%d]", p2, p1)
	if s1 == "" {
		return aString
	}
	g := strings.Split(s2, s1+"]")
	left := addToLastNumber(g[0], p1)
	right := addToNextNumber(g[1], p2)
	// left := addToLastNumber(aString[:ix], p1)
	// right := addToNextNumber(aString[ix+len(pr):], p2)
	out := left + "0" + right
	return out
}

func addToNextNumber(aString string, x int) string {
	// var num int
	// nextNumberScan := "%d"
	ns := getFirstNumber(aString)
	if ns == "-1" {
		return aString
	}
	nbr, _ := strconv.Atoi(ns)
	insertion := fmt.Sprint(nbr + x)
	newString := strings.Replace(aString, ns, insertion, 1)
	return newString
	// for i := range aString {
	// 	s := aString[i:]
	// 	n, _ := fmt.Sscanf(s, nextNumberScan, &num)
	// 	if n == 1 {
	// 		break
	// 	}
	// }
	// newString := strings.Replace(aString, fmt.Sprint(num), fmt.Sprint(num+x), 1)
	// return newString
}

func getFirstNumber(aString string) string {
	numbers := []string{}
	accum := ""
	for _, v := range aString {
		c := string(v)
		if isDigit(c) {
			accum += c
		} else {
			if len(accum) > 0 {
				numbers = append(numbers, accum)
				accum = ""
			}
		}

	}
	if len(accum) > 0 {
		numbers = append(numbers, accum)
	}

	if len(numbers) > 0 {
		return numbers[0]
	}
	return "-1"
}

func getLastNumber(aString string) string {
	numbers := []string{}
	accum := ""
	for _, v := range aString {
		c := string(v)
		if isDigit(c) {
			accum += c
		} else {
			if len(accum) > 0 {
				numbers = append(numbers, accum)
				accum = ""
			}
		}

	}
	if len(accum) > 0 {
		numbers = append(numbers, accum)
	}
	if len(numbers) > 0 {
		return numbers[len(numbers)-1]
	}
	return "-1"
}

func addToLastNumber(aString string, x int) string {
	// var num int
	// newRep := ""
	// lastNumberScan := "%d"
	// get the last number in the string
	oldNum := getLastNumber(aString)
	if oldNum == "-1" {
		return aString
	}
	newValue, _ := strconv.Atoi(oldNum)
	newValue += x
	sNewValue := fmt.Sprint(newValue)
	rsNewValue := reverse(sNewValue)

	roldNum := reverse(oldNum)

	reversed := reverse(aString)
	s := strings.Replace(reversed, roldNum, rsNewValue, 1)
	return reverse(s)

	// for i := range r {
	// 	s := r[i:]
	// 	n, _ := fmt.Sscanf(s, lastNumberScan, &num)

	// 	if n == 1 {
	// 		lastNumberString := reverse(fmt.Sprint(num))
	// 		oldNum, _ := strconv.Atoi(lastNumberString)
	// 		newNumber := oldNum + x
	// 		newRep = reverse(fmt.Sprint(newNumber))
	// 		break
	// 	}

	// r := reverse(aString)
	// for i := range r {
	// 	s := r[i:]
	// 	n, _ := fmt.Sscanf(s, lastNumberScan, &num)
	// 	if n == 1 {
	// 		lastNumberString := reverse(fmt.Sprint(num))
	// 		oldNum, _ := strconv.Atoi(lastNumberString)
	// 		newNumber := oldNum + x
	// 		newRep = reverse(fmt.Sprint(newNumber))
	// 		break
	// 	}
	// }
	// return reverse(strings.Replace(r, fmt.Sprint(num), newRep, 1))

}

func isDigit(aString string) bool {
	return len(aString) == 1 && strings.Contains("0123456789", aString)
}

func reverse(aString string) string {
	c := len(aString)
	sList := ""
	for v := range aString {
		sList += string(aString[c-1-v])
	}
	return sList
}

func digitOver10(aString string) int {

	accum := ""
	for _, v := range aString {
		if isDigit(string(v)) {
			accum += string(v)
			val, _ := strconv.Atoi(accum)
			if val >= 10 {
				return val
			}
		} else {
			accum = ""
		}
	}
	return 0
}

func digitSplit(anumber int) string {
	first := anumber / 2
	second := anumber/2 + anumber%2
	return fmt.Sprintf("[%d,%d]", first, second)
}
