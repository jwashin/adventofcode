package main

import "fmt"

func main() {
	// fmt.Println("1.", firstPart())
	fmt.Println("2.", part2())
}

// TODO: optimize for speed. probably dynamic programming

func part2() int {
	s := "1113122113"
	for i := 0; i < 50; i++ {
		s = lookandsay(s)
	}
	return len(s)
}

func firstPart() int {
	s := "1113122113"
	for i := 0; i < 40; i++ {
		s = lookandsay(s)
	}
	return len(s)

}

func lookandsay(s string) string {
	countString := ""
	outString := ""

	for k, v := range s {
		if byte(v) != s[0] {
			outString = fmt.Sprintf("%d%s", len(countString), string(countString[0]))
			return outString + lookandsay(s[k:])
		}
		countString += string(v)
	}
	return fmt.Sprintf("%d%s", len(countString), string(countString[0]))
}
