package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println("1: " + allCodes("5", string(data)))
}

var dests = map[string]string{"U": "123123456", "D": "456789789", "L": "112445778", "R": "233566899"}

var keypad = "123456789"

func nextButton(currentButton string, instruction string) string {

	currentIndex := strings.Index(keypad, currentButton)
	i := dests[instruction]
	c := i[currentIndex]
	return string(c)
}

func nextCode(start string, instructions string) string {
	cb := start
	for _, v := range instructions {
		i := string(v)
		cb = nextButton(cb, i)
	}
	return cb

}

func allCodes(start string, data string) string {
	cb := start
	output := ""
	parsed := strings.Split(data, "\n")
	for _, value := range parsed {
		d := strings.TrimSpace(value)
		if len(d) > 0 {
			res := nextCode(cb, d)
			output += res
			cb = res
		}
	}
	return output
}
