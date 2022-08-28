package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func main() {
	// salt := "zpqevtbw"
	// fmt.Printf("1. %d", makePad(salt))
	// salt := "abc"
	// fmt.Printf("2demo. %d", superMakePad(salt))
	salt := "zpqevtbw"
	fmt.Printf("2. %d", superMakePad(salt))
}

func calcMD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func superHash(s string) string {
	idx := 0
	for idx < 2017 {
		idx += 1
		s = calcMD5(s)
	}
	return s
}

func findFirstTriplet(aString string) string {
	s := calcMD5(aString)
	r, _ := regexp.Compile(`(aaa|bbb|ccc|ddd|eee|fff|000|111|222|333|444|555|666|777|888|999.*?)`)
	d := r.FindString(s)
	if d != "" {
		return string(d[0])
	}
	return ""
}

func findSuperFirstTriplet(aString string) string {
	s := superHash(aString)
	r, _ := regexp.Compile(`(aaa|bbb|ccc|ddd|eee|fff|000|111|222|333|444|555|666|777|888|999.*?)`)
	d := r.FindString(s)
	if d != "" {
		return string(d[0])
	}
	return ""
}

func makePad(salt string) int {
	md5s := map[int]string{}
	pad := []int{}
	idx := 0
	for len(pad) < 64 {
		idx += 1
		currRepeatChar := findFirstTriplet(fmt.Sprintf("%s%d", salt, idx))
		if currRepeatChar != "" {
			checkIdx := idx + 1000
			ddx := idx
			// it := strings.Repeat(currRepeatChar, 5)
			r, _ := regexp.Compile(`.*` + currRepeatChar + `{5}.*`)
			for ddx <= checkIdx {
				ddx += 1
				ky := fmt.Sprintf("%s%d", salt, ddx)
				s := md5s[ddx]
				if s == "" {
					s = calcMD5(ky)
					md5s[ddx] = s
				}
				if r.FindString(s) != "" {
					pad = append(pad, idx)
					fmt.Printf("%d. for  %d, '%s' (%s)  '%s' at index %d (%d)\n", len(pad), idx, currRepeatChar, md5s[idx], s, ddx, ddx-idx)

					continue
				}
			}
		}
	}
	return idx

}

func superMakePad(salt string) int {
	md5s := map[int]string{}
	pad := []int{}
	idx := 0
	for len(pad) < 64 {
		idx += 1
		currRepeatChar := findSuperFirstTriplet(fmt.Sprintf("%s%d", salt, idx))
		if currRepeatChar != "" {
			checkIdx := idx + 1000
			ddx := idx
			// it := strings.Repeat(currRepeatChar, 5)
			r, _ := regexp.Compile(`.*` + currRepeatChar + `{5}.*`)
			for ddx <= checkIdx {
				ddx += 1
				ky := fmt.Sprintf("%s%d", salt, ddx)
				s := md5s[ddx]
				if s == "" {
					s = superHash(ky)
					md5s[ddx] = s
				}
				if r.FindString(s) != "" {
					pad = append(pad, idx)
					fmt.Printf("%d. for  %d, '%s' (%s)  '%s' at index %d (%d)\n", len(pad), idx, currRepeatChar, md5s[idx], s, ddx, ddx-idx)

					continue
				}
			}
		}
	}
	return idx

}
