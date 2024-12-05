package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1:", part1(string(input)))
	fmt.Println("part 2:", part2(string(input)))
}

func getData(s string) (rules []string, updates []string) {

	input := strings.Split(s, "\n")
	rules = []string{}
	updates = []string{}

	doRules := true

	for _, v := range input {
		if len(v) == 0 {
			doRules = false
			continue
		}
		if doRules {
			rules = append(rules, v)
		} else {
			updates = append(updates, v)
		}
	}
	return rules, updates
}

func ruleApplies(rule []string, update string) bool {
	for _, v := range rule {
		if !strings.Contains(update, v) {
			return false
		}
	}
	return true
}

func isCorrect(update string, rules []string) bool {
	for _, v := range rules {
		t := strings.Split(v, "|")
		if ruleApplies(t, update) {
			if strings.Index(update, t[0]) > strings.Index(update, t[1]) {
				return false
			}
		}
	}
	return true
}

func part1(s string) int {
	rules, updates := getData(s)
	result := 0
	for _, update := range updates {
		if isCorrect(update, rules) {
			items := strings.Split(update, ",")
			middle := items[len(items)/2]
			pageNumber, _ := strconv.Atoi(middle)
			result += pageNumber
		}
	}
	return result
}

func part2(s string) int {
	rules, updates := getData(s)
	result := 0
	for _, update := range updates {
		if !isCorrect(update, rules) {

			for !isCorrect(update, rules) {
				update = reorderUpdate(update, rules)
			}
			items := strings.Split(update, ",")
			middle := items[len(items)/2]
			pageNumber, _ := strconv.Atoi(middle)
			result += pageNumber
		}
	}
	return result

}

func reorderUpdate(update string, rules []string) string {
	for _, v := range rules {
		t := strings.Split(v, "|")
		if ruleApplies(t, update) {
			if strings.Index(update, t[0]) > strings.Index(update, t[1]) {
				// swap t0 and t1 in the update
				t0 := t[0]
				t1 := t[1]
				tmp := "mmmm"
				update = strings.Replace(update, t0, tmp, 1)
				update = strings.Replace(update, t1, t0, 1)
				update = strings.Replace(update, tmp, t1, 1)
			}
		}
	}
	return update
}
