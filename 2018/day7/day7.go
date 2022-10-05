package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

func part1() string {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	r := orderSteps(data)
	return r
}

func part2() int {
	input, _ := os.ReadFile("input.txt")
	data := string(input)
	r := makeSleigh(data, 5, false)
	return r
}

type cond struct {
	precedent string
	step      string
}

type rules []cond

func (r rules) allComplete(s string) bool {
	availableSteps := map[string]bool{}
	for _, v := range r {
		availableSteps[v.precedent] = true
		availableSteps[v.step] = true
	}
	for _, v := range s {
		delete(availableSteps, string(v))
	}
	return len(availableSteps) == 0
}

func (r rules) chooseNext(s string, availableWorkers int) []string {
	completed := map[string]bool{}
	for _, v := range s {
		completed[string(v)] = true
	}
	availableSteps := map[string]bool{}
	for _, v := range r {
		if !completed[v.precedent] {
			availableSteps[v.precedent] = true
		}
		if !completed[v.step] {
			availableSteps[v.step] = true
		}
	}
	for _, v := range r {
		if !completed[v.precedent] {
			delete(availableSteps, v.step)
		}
	}
	items := []string{}
	for v := range availableSteps {
		items = append(items, v)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	if availableWorkers < len(items) {
		return items[:availableWorkers]
	}
	return items
}

func makeSteps(s string) rules {
	steps := rules{}
	for _, v := range strings.Split(s, "\n") {
		t := strings.Fields(v)
		f := cond{step: t[7], precedent: t[1]}
		steps = append(steps, f)
	}
	return steps
}

func orderSteps(s string) string {
	r := makeSteps(s)
	out := r.chooseNext("", 1)
	path := ""
	for len(out) > 0 {
		path += out[0]
		out = r.chooseNext(path, 1)
	}
	return path
}

type task struct {
	timeRemaining int
}

func (t *task) dec() {
	t.timeRemaining -= 1
}

func makeSleigh(s string, elves int, test bool) int {
	alpha := "0ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r := makeSteps(s)
	// out := r.chooseNext("")
	path := ""
	time := -1
	availableElves := elves
	tasks := map[string]*task{}
	newTasks := r.chooseNext(path, availableElves)
	// newTasks := []string{}
	for !r.allComplete(path) {
		getNewTasks := false
		for k, v := range tasks {
			if v.timeRemaining == 0 {
				delete(tasks, k)
				availableElves += 1
				path += k
				getNewTasks = true
			}
		}
		if getNewTasks {
			newTasks = r.chooseNext(path, availableElves)
		}
		for _, newTask := range newTasks {
			// chooseNext is based on completed items, so don't
			// restart any in-progress
			found := false
			for k := range tasks {
				if k == newTask {
					found = true
				}
			}
			if !found {
				taskDuration := 60 + strings.Index(alpha, newTask)
				if test {
					taskDuration -= 60
				}
				if availableElves > 0 {
					tasks[newTask] = &task{taskDuration}
					availableElves -= 1
				}
			}
		}
		newTasks = []string{}
		time += 1
		for _, v := range tasks {
			v.dec()
		}
	}
	return time
}
