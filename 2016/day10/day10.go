package main

import (
	"fmt"
	"strings"
)

func whotests(instructions string, compa int, compb int) {
	containers := map[string][]int{}
	bots := map[string]map[string]string{}
	for _, instruction := range strings.Split(instructions, "\n") {
		if strings.Contains(instruction, "value") {
			var value int
			var botNumber string
			fmt.Sscanf(instruction, "value %d goes to %s", &value, &botNumber)
			containers["bot "+botNumber] = append(containers["bot "+botNumber], value)
		}
		if strings.Contains(instruction, "gives") {
			var botNumber, first, second string
			fmt.Sscanf(instruction, "bot %s gives low to %d and high to %d", &botNumber, &first, &second)
			bots["bot "+botNumber] = map[string]string{"high": second, "low": first}
		}
	}
}
