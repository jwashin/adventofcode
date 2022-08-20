package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println("1: " + allCodes("5", string(data)))
	fmt.Println("2: " + allCodes2("5", string(data)))
}

func nextButton(currentButton string, instruction string) string {
	keypad := "123456789"
	dests := map[string]string{"U": "123123456", "D": "456789789", "L": "112445778", "R": "233566899"}
	currentIndex := strings.Index(keypad, currentButton)
	i := dests[instruction]
	c := i[currentIndex]
	return string(c)
}

func nextButton2(currentButton string, instruction string) string {
	keypad := "123456789ABCD"
	dests := map[string]string{"U": "121452349678B", "D": "36785ABC9ADCD",
		"L": "122355678AABD", "R": "134467899BCCD"}
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

func nextCode2(start string, instructions string) string {
	cb := start
	for _, v := range instructions {
		i := string(v)
		cb = nextButton2(cb, i)
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
func allCodes2(start string, data string) string {
	cb := start
	output := ""
	parsed := strings.Split(data, "\n")
	for _, value := range parsed {
		d := strings.TrimSpace(value)
		if len(d) > 0 {
			res := nextCode2(cb, d)
			output += res
			cb = res
		}
	}
	return output
}
