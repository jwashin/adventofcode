package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s, _ := os.ReadFile("input.txt")
	data := string(s)
	fmt.Println("1.", lowestValAllowed(data))
	fmt.Println("2.", countAllowed(data))
}

type interval struct {
	min int
	max int
}

type Intervals []*interval

func (e Intervals) Len() int {
	return len(e)
}
func (a Intervals) Less(i, j int) bool { return a[i].min < a[j].min }
func (a Intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a Intervals) isAllowed(t int) (bool, *interval) {
	for _, v := range a {
		if v.contains(t) {
			return false, v
		}
	}
	return true, &interval{0, 0}
}
func (a Intervals) contains(num int) bool {
	for _, v := range a {
		if v.contains(num) {
			return true
		}
	}
	return false
}

func (a Intervals) add(candidate interval) Intervals {
	sort.Sort(a)
	for _, item := range a {

		if item.contains(candidate.max) && item.contains(candidate.min) {
			return a
		}
		if candidate.contains(item.max) && candidate.contains(item.min) {
			item.min = candidate.min
			item.max = candidate.max
			return a
		}

		if item.contains(candidate.max) {
			if candidate.min < item.min {
				item.min = candidate.min
			}
			return a
		}
		if item.contains(candidate.min) {
			if candidate.max > item.max {
				item.max = candidate.max
			}
			return a
		}
	}
	a = append(a, &candidate)
	a = a.normalize()
	sort.Sort(a)
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a Intervals) normalize() Intervals {
	sort.Sort(a)
	lastIndex := len(a) - 1
	s := Intervals{}
	for idx, val := range a {
		if idx == lastIndex {
			if val.max != val.min {
				s = append(s, val)
			}
			break
		}
		if a[idx+1].contains(val.max) || val.contains(a[idx+1].min) || a[idx+1].min == val.max+1 {
			val.max = max(a[idx+1].max, val.max)
			val.min = min(a[idx+1].min, val.min)
			a[idx+1].min = a[idx+1].max
			// s = append(s, val)
		}

		if val.max != val.min {
			newItem := interval{}
			newItem.set(fmt.Sprintf("%d-%d", val.max, val.min))
			s = append(s, &newItem)

		}

	}
	sort.Sort(s)
	return s
}

func (i *interval) set(s string) {
	t := strings.Split(strings.TrimSpace(s), "-")
	ints := []int{}
	for _, v := range t {
		a, _ := strconv.Atoi(v)
		ints = append(ints, a)
	}
	if len(ints) == 2 {
		if ints[0] < ints[1] {
			i.min = ints[0]
			i.max = ints[1]
			return
		}
		i.min = ints[1]
		i.max = ints[0]
	}
}

func (i interval) contains(a int) bool {
	if a >= i.min && a <= i.max {
		return true
	}
	return false
}

func makeIntervals(s string) Intervals {
	intervals := Intervals{}
	t := strings.Split(s, "\n")
	for _, v := range t {
		ni := interval{}
		ni.set(v)
		intervals = intervals.add(ni)
	}
	return intervals
}

func lowestValAllowed(s string) int {
	intervals := Intervals(makeIntervals(s))
	sort.Sort(intervals)
	newIntervalLen := 100
	for len(intervals) != newIntervalLen {
		newIntervalLen = len(intervals)
		intervals = intervals.normalize()
	}
	min := -1
	allowed := false
	for !allowed {
		min += 1
		allowed, interval := intervals.isAllowed(min)
		if allowed {
			break
		}
		min = interval.max
	}
	return min
}

// 1468503307 too high
// 681316812 too high
// 682138306 too high
// 891587016 too high
// 977912121 too high
// 287153283 too high
//  also wrong 287153282
// 173040686
// 84028932 also wrong

func countAllowed(s string) int {
	intervals := Intervals(makeIntervals(s))

	// brute force!
	count := 0
	num := 0
	for num <= math.MaxUint32 {
		if !intervals.contains(num) {
			count += 1
		}
		num += 1
		if num%10000000 == 0 {
			fmt.Println(num, count)
		}
	}

	// 3830000000
	// sort.Sort(intervals)
	// newIntervalLen := 100
	// for len(intervals) != newIntervalLen {
	// 	newIntervalLen = len(intervals)
	// 	intervals = intervals.normalize()
	// }
	// overmax := math.MaxUint32
	// whitelistCount := intervals[0].min
	// max := intervals[0].max
	// for idx, v := range intervals {
	// 	if idx == 0 {
	// 		continue
	// 	}
	// 	whitelistCount += v.min - max - 1
	// 	max = v.max
	// }
	// whitelistCount += (overmax - max)
	return count
}
