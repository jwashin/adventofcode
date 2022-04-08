package lanternfish

import (
	"strconv"
	"strings"
)

type Lanternfish struct {
	timer int
}

func CountFish(aString string, days int) int {
	school := make(map[int]int)
	alfas := strings.Split(strings.TrimSpace(aString), ",")
	for _, val := range alfas {
		timer, _ := strconv.Atoi(val)
		school[timer] += 1
	}

	for day := 1; day <= days; day++ {
		newSchool := make(map[int]int)
		for timer, count := range school {
			if timer == 0 {
				newSchool[6] += count
				newSchool[8] += count
			} else {
				newSchool[timer-1] += count
			}
		}
		school = newSchool
	}
	count := 0
	for _, fish := range school {
		count += fish
	}
	return count
}

func OldCountFish(aString string, days int) int {
	school := []Lanternfish{}
	alfas := strings.Split(strings.TrimSpace(aString), ",")
	for _, val := range alfas {
		timer, _ := strconv.Atoi(val)
		school = append(school, Lanternfish{timer})
	}

	for day := 0; day < days; day++ {
		newSchool := []Lanternfish{}
		for _, fish := range school {
			if fish.timer == 0 {
				newSchool = append(newSchool, Lanternfish{6})
				newSchool = append(newSchool, Lanternfish{8})
			} else {
				newSchool = append(newSchool, Lanternfish{fish.timer - 1})
			}
		}
		school = newSchool
	}
	return len(school)
}
