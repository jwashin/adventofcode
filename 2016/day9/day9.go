package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	j := decompress(string(input))
	fmt.Println(len(j))
	fmt.Println(decompressedLength(j))
}

func decompressedLength(s string) int {
	out := 0
	gettingMarker := false
	marker := ""
	repeatString := ""
	gettingRepeatString := false
	length := 0
	count := 0
	for _, v := range s {
		if gettingMarker {
			if string(v) == ")" {
				// close command, parse and start acquiring n items for expansion
				// var length, count int
				fmt.Sscanf(marker, "%dx%d", &length, &count)
				repeatString = ""
				gettingMarker = false
				gettingRepeatString = true
				continue
			}
			// not closing the marker yet, so add to cmd string
			marker += string(v)
			continue
		}
		if gettingRepeatString && len(repeatString) < length {
			repeatString += string(v)
			if len(repeatString) == length {
				out += decompressedLength(repeatString) * count
				// for count > 0 {
				// 	accum += repeatString
				// 	count -= 1
				// }
				marker = ""
			}
			continue
		}
		if string(v) == "(" {
			gettingMarker = true
			continue
		}
		out += 1
	}
	return out
}

func decompress(s string) string {
	accum := ""
	gettingMarker := false
	marker := ""
	repeatString := ""
	gettingRepeatString := false
	length := 0
	count := 0
	for _, v := range s {
		if gettingMarker {
			if string(v) == ")" {
				// close command, parse and start acquiring n items for expansion
				// var length, count int
				fmt.Sscanf(marker, "%dx%d", &length, &count)
				repeatString = ""
				gettingMarker = false
				gettingRepeatString = true
				continue
			}
			// not closing the marker yet, so add to cmd string
			marker += string(v)
			continue
		}
		if gettingRepeatString && len(repeatString) < length {
			repeatString += string(v)
			if len(repeatString) == length {
				for count > 0 {
					accum += repeatString
					count -= 1
				}
				marker = ""
			}
			continue
		}
		if string(v) == "(" {
			gettingMarker = true
			continue
		}
		accum += string(v)
	}
	return accum
}
