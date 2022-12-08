package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("Part 1", part1(string(input)))
	fmt.Println("Part 2", part2(string(input)))
}

type FSDirectory struct {
	parent  *FSDirectory
	name    string
	files   []*FSFile
	subdirs []*FSDirectory
}
type FSFile struct {
	name string
	size int
}

func (f *FSDirectory) size() int {
	total := 0
	for _, v := range f.files {
		total += v.size
	}
	for _, v := range f.subdirs {
		total += v.size()
	}
	return total
}

func (f *FSDirectory) allSubdirectories() []*FSDirectory {
	theList := []*FSDirectory{}
	theList = append(theList, f.subdirs...)
	for _, v := range f.subdirs {
		theList = append(theList, v.allSubdirectories()...)
	}
	return theList
}

func (f *FSDirectory) chdir(s string) (*FSDirectory, error) {
	if s == "/" {
		for f.parent != nil {
			f = f.parent
		}
		return f, nil
	}
	if s == ".." {
		return f.parent, nil
	}
	for _, v := range f.subdirs {
		if v.name == s {
			return v, nil
		}
	}
	return nil, errors.New("No such directory: " + s)
}

func part1(s string) int {
	fs := dirMaker(s)
	theList := fs.allSubdirectories()
	total := 0
	for _, v := range theList {
		if v.size() <= 100000 {
			total += v.size()
		}
	}
	return total
}

// 1432936 too high OOps, needed to actually invoke part2 in main()
func part2(s string) int {
	totalSpace := 70000000
	needed := 30000000
	fs := dirMaker(s)
	used := fs.size()
	available := totalSpace - used

	target := needed - available

	theList := fs.allSubdirectories()
	cands := []int{}
	for _, v := range theList {
		sz := v.size()
		if sz >= target {
			cands = append(cands, sz)
		}
	}
	min := cands[0]
	for _, v := range cands {
		if v < min {
			min = v
		}
	}
	return min
}

func dirMaker(s string) *FSDirectory {

	root := FSDirectory{name: "/"}

	currLoc := &root

	input := strings.Split(s, "\n")

	for _, v := range input {
		f := strings.Fields(v)
		if f[0] == "$" {
			if f[1] == "cd" {
				currLoc, _ = currLoc.chdir(f[2])
			}
			continue
		}
		if f[0] == "dir" {
			newDir := FSDirectory{name: f[1], parent: currLoc}
			currLoc.subdirs = append(currLoc.subdirs, &newDir)
			continue
		}
		z, err := strconv.Atoi(f[0])
		if err == nil {
			newItem := FSFile{name: f[1], size: z}
			currLoc.files = append(currLoc.files, &newItem)
		}
	}
	return &root
}
